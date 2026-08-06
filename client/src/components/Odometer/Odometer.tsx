import { useEffect, useRef, useState } from "react";
import "./Odometer.scss";
import { ANIMATION_SPEED, ODOMETER } from "../../config/constants";
import { OdometerProps } from "./Odometer.types";

/**
 * Scale markers, laid out clockwise from the top so they follow the direction
 * of the arc sweep
 */
const TICKS = [
  { value: 20, position: "counter1" },
  { value: 70, position: "counter2" },
  { value: 100, position: "counter3" },
  { value: 150, position: "counter4" },
];

/**
 * Odometer component
 * Displays vehicle speed as a numeric readout plus a circular arc gauge.
 * Purely presentational: the caller owns the store subscription
 * @param {OdometerProps} props - Current speed and the range it is drawn against
 */
const Odometer = ({ value, min, max, className }: OdometerProps) => {
  const [speed, setSpeed] = useState(0);
  const requestRef = useRef<number | null>(null);
  const [isRaspberryPi, setIsRaspberryPi] = useState(false);

  /**
   * Anima gradualmente il valore velocità verso il target
   */
  const animateSpeed = () => {
    setSpeed((prev) => {
      const diff = value - prev;
      const step = Math.sign(diff) * Math.min(Math.abs(diff), ANIMATION_SPEED.STEP);
      const next = prev + step;

      if (Math.abs(diff) <= ANIMATION_SPEED.THRESHOLD) {
        cancelAnimationFrame(requestRef.current!);
        return value;
      }

      requestRef.current = requestAnimationFrame(animateSpeed);
      return next;
    });
  };

  /**
   * Rileva se l'applicazione è in esecuzione su Raspberry Pi
   */
  useEffect(() => {
    const userAgent = navigator.userAgent.toLowerCase();
    const platform = navigator.platform.toLowerCase();
    const isRpi = userAgent.includes('linux arm') ||
                  platform.includes('arm') ||
                  userAgent.includes('raspberry') ||
                  (navigator.hardwareConcurrency !== undefined && navigator.hardwareConcurrency <= 4);

    setIsRaspberryPi(isRpi);
  }, []);

  /**
   * Aggiorna il valore velocità quando cambia lo stato
   */
  useEffect(() => {
    if (isRaspberryPi) {
      setSpeed(value);
    } else {
      requestRef.current = requestAnimationFrame(animateSpeed);
    }

    return () => {
      if (requestRef.current) cancelAnimationFrame(requestRef.current);
    };
  }, [value, isRaspberryPi]);

  const range = Math.max(max - min, 1);
  const progress = Math.min(Math.max((speed - min) / range, 0), 1);
  const percentage = progress * ODOMETER.ARC_SWEEP;

  /**
   * Determina il colore in base alla velocità
   */
  const getColor = () => {
    if (progress < 0.7) return "rgba(123, 212, 211, 0.6)";
    if (progress < 0.9) return "rgba(255, 255, 255, 0.6)";
    return "rgba(255, 0, 0, 0.6)";
  };

  return (
    <div className={className ? `componentOdometer ${className}` : "componentOdometer"}>
      <div className="wrapper">
        <div
          className="circle"
          style={{
            background: `conic-gradient(${getColor()} 0% ${percentage}%, transparent ${percentage}% 100%)`,
          }}
        >
          {TICKS.map(({ value: tick, position }) => (
            <span key={tick} className={position}>
              {tick}
            </span>
          ))}

          <div className="inner" />
          <div className="mid" />
          <div className="label">
            <h2>{Math.round(speed)}</h2>
            <p>km/h</p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Odometer;
