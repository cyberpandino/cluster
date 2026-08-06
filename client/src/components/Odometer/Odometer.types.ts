/**
 * Props of the Odometer component
 * The range is passed in rather than read from the store so the scale markers
 * always match the gauge the caller is actually driving
 */
export interface OdometerProps {
  /** Current speed in km/h. Values outside the range are clamped when drawn */
  value: number;
  /** Lower bound of the scale in km/h */
  min: number;
  /** Upper bound of the scale in km/h, and the value of the last tick */
  max: number;
  /** Escape hatch for one-off layout overrides */
  className?: string;
}
