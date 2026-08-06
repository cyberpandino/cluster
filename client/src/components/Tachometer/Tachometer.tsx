import { useSnapshot } from "valtio";
import Gauge from "../Gauge";
import { state } from "../../store/state";

const TICKS = [1, 3, 5, 7] as const;

/**
 * Componente Tachometer
 * Visualizza i giri motore (RPM)
 */
const Tachometer = () => {
  const rpm = useSnapshot(state.session.rpm);

  return <Gauge value={rpm.current} min={rpm.min} max={rpm.max} ticks={TICKS} unit="RPM" />;
};

export default Tachometer;
