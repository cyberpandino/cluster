import { CSSProperties, useEffect, useRef, useState } from "react";
import styles from "./Gauge.module.scss";
import { GAUGE } from "../../config/constants";
import { graphics } from "../../config/environment";
import { GaugeProps } from "./Gauge.types";

const REDUCED_MOTION_QUERY = "(prefers-reduced-motion: reduce)";

/**
 * Corner slots for the scale markers, clockwise from the top so they follow
 * the direction of the arc sweep
 */
const TICK_POSITIONS = [
  styles.tickTopRight,
  styles.tickBottomRight,
  styles.tickBottomLeft,
  styles.tickTopLeft,
];

/**
 * Maps progress to a severity level, exposed as a data attribute so the
 * stylesheet owns the arc colours instead of splitting them across TS and SCSS
 * @param {number} progress - Reading as a fraction of the range, from 0 to 1
 * @returns {string} Level consumed by the [data-level] selectors
 */
const getLevel = (progress: number) => {
  if (progress < GAUGE.HIGH_LEVEL) return "normal";
  if (progress < GAUGE.CRITICAL_LEVEL) return "high";
  return "critical";
};

/**
 * Gauge component
 * Circular dial with an arc, four corner markers and a numeric readout.
 * Purely presentational: the caller owns the store subscription
 * @param {GaugeProps} props - Current reading and the scale it is drawn against
 */
const Gauge = ({ value, min, max, ticks, unit, className }: GaugeProps) => {
  const [displayed, setDisplayed] = useState(value);
  const [reducedMotion, setReducedMotion] = useState(
    () => window.matchMedia(REDUCED_MOTION_QUERY).matches
  );

  const targetRef = useRef(value);
  const animatedRef = useRef(value);
  const frameRef = useRef<number | null>(null);

  const isInstant = reducedMotion || graphics.quality === 1;

  /**
   * Tracks the accessibility preference while mounted, since the cluster runs
   * for hours and the setting can change without a reload
   */
  useEffect(() => {
    const query = window.matchMedia(REDUCED_MOTION_QUERY);
    const handleChange = (event: MediaQueryListEvent) => setReducedMotion(event.matches);

    query.addEventListener("change", handleChange);
    return () => query.removeEventListener("change", handleChange);
  }, []);

  /**
   * Drives the displayed value towards the incoming reading.
   * A single loop is kept alive across value changes: restarting it on every
   * update would reset the easing and make the readout stutter
   */
  useEffect(() => {
    targetRef.current = value;

    if (isInstant) {
      if (frameRef.current !== null) cancelAnimationFrame(frameRef.current);
      frameRef.current = null;
      animatedRef.current = value;
      setDisplayed(value);
      return;
    }

    if (frameRef.current !== null) return;

    let previousTime = performance.now();

    // Exponential smoothing: frame-rate independent, and able to retarget
    // without restarting the animation
    const tick = (time: number) => {
      const delta = time - previousTime;
      previousTime = time;

      const remaining = targetRef.current - animatedRef.current;

      if (Math.abs(remaining) < GAUGE.SETTLE_THRESHOLD) {
        animatedRef.current = targetRef.current;
        frameRef.current = null;
        setDisplayed(targetRef.current);
        return;
      }

      animatedRef.current += remaining * (1 - Math.exp(-delta / GAUGE.SMOOTHING));
      frameRef.current = requestAnimationFrame(tick);
      setDisplayed(animatedRef.current);
    };

    frameRef.current = requestAnimationFrame(tick);
  }, [value, isInstant]);

  /**
   * Unmount-only cleanup. It cannot live in the effect above, whose cleanup
   * would cancel the running loop on every value update
   */
  useEffect(() => {
    return () => {
      if (frameRef.current !== null) cancelAnimationFrame(frameRef.current);
      frameRef.current = null;
    };
  }, []);

  const range = Math.max(max - min, 1);
  const progress = Math.min(Math.max((displayed - min) / range, 0), 1);

  return (
    <div className={className ? `${styles.root} ${className}` : styles.root}>
      <div className={styles.wrapper}>
        <div
          className={styles.dial}
          data-level={getLevel(progress)}
          style={{ "--gauge-fill": `${progress * GAUGE.ARC_SWEEP}%` } as CSSProperties}
        >
          {ticks.map((tick, index) => (
            <span key={tick} className={`${styles.tick} ${TICK_POSITIONS[index]}`}>
              {tick}
            </span>
          ))}

          <div className={styles.ring} />
          <div className={styles.hub} />
          <div className={styles.readout}>
            <h2>{Math.round(displayed)}</h2>
            <p>{unit}</p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Gauge;
