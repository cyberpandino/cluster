import { CSSProperties, useEffect, useRef, useState } from "react";
import styles from "./Odometer.module.scss";
import { ODOMETER } from "../../config/constants";
import { graphics } from "../../config/environment";
import { OdometerProps } from "./Odometer.types";

const REDUCED_MOTION_QUERY = "(prefers-reduced-motion: reduce)";

/**
 * Scale markers, laid out clockwise from the top so they follow the direction
 * of the arc sweep
 */
const TICKS = [
  { value: 20, position: styles.tickTopRight },
  { value: 70, position: styles.tickBottomRight },
  { value: 100, position: styles.tickBottomLeft },
  { value: 150, position: styles.tickTopLeft },
];

/**
 * Maps progress to a severity level, exposed as a data attribute so the
 * stylesheet owns the arc colours instead of splitting them across TS and SCSS
 * @param {number} progress - Speed as a fraction of the range, from 0 to 1
 * @returns {string} Level consumed by the [data-level] selectors
 */
const getLevel = (progress: number) => {
  if (progress < ODOMETER.HIGH_LEVEL) return "normal";
  if (progress < ODOMETER.CRITICAL_LEVEL) return "high";
  return "critical";
};

/**
 * Odometer component
 * Displays vehicle speed as a numeric readout plus a circular arc gauge.
 * Purely presentational: the caller owns the store subscription
 * @param {OdometerProps} props - Current speed and the range it is drawn against
 */
const Odometer = ({ value, min, max, className }: OdometerProps) => {
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
   * Drives the displayed value towards the incoming speed.
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

      if (Math.abs(remaining) < ODOMETER.SETTLE_THRESHOLD) {
        animatedRef.current = targetRef.current;
        frameRef.current = null;
        setDisplayed(targetRef.current);
        return;
      }

      animatedRef.current += remaining * (1 - Math.exp(-delta / ODOMETER.SMOOTHING));
      frameRef.current = requestAnimationFrame(tick);
      setDisplayed(animatedRef.current);
    };

    frameRef.current = requestAnimationFrame(tick);
  }, [value, isInstant]);

  /**
   * Unmount-only cleanup. It cannot live in the effect above, whose cleanup
   * would cancel the running loop on every speed update
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
          style={{ "--odometer-fill": `${progress * ODOMETER.ARC_SWEEP}%` } as CSSProperties}
        >
          {TICKS.map(({ value: tick, position }) => (
            <span key={tick} className={`${styles.tick} ${position}`}>
              {tick}
            </span>
          ))}

          <div className={styles.ring} />
          <div className={styles.hub} />
          <div className={styles.readout}>
            <h2>{Math.round(displayed)}</h2>
            <p>km/h</p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Odometer;
