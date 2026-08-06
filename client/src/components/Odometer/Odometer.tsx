import { useSnapshot } from "valtio";
import Gauge from "../Gauge";
import { state } from "../../store/state";

const TICKS = [20, 70, 100, 150] as const;

/**
 * Componente Odometer
 * Visualizza la velocità del veicolo
 */
const Odometer = () => {
  const speed = useSnapshot(state.session.speed);

  return <Gauge value={speed.current} min={speed.min} max={speed.max} ticks={TICKS} unit="km/h" />;
};

export default Odometer;
