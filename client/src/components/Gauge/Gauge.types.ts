/**
 * Props of the Gauge component
 * The range is passed in rather than read from the store, so the component
 * stays presentational and the caller owns the subscription
 */
export interface GaugeProps {
  /** Current reading. Values outside the range are clamped when drawn */
  value: number;
  /** Lower bound of the scale */
  min: number;
  /** Upper bound of the scale */
  max: number;
  /**
   * Labels for the four corner markers, clockwise from the top.
   * The dial has exactly four slots, so the tuple length is enforced
   */
  ticks: readonly [number, number, number, number];
  /** Unit shown under the reading, for example "km/h" or "RPM" */
  unit: string;
  /** Escape hatch for one-off layout overrides */
  className?: string;
}
