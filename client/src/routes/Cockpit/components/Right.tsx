import Odometer from "../../../components/Odometer";
import Battery from "../../../components/Battery";
import { useSnapshot } from "valtio";
import { state } from "../../../store/state";


const Right = () => {
  const snapSpeed = useSnapshot(state.session.speed);

  return (
    <div className="right">
        <div className="right__center">
            <Odometer value={snapSpeed.current} min={snapSpeed.min} max={snapSpeed.max} />
        </div>
        <div className="right__bottom">
          <Battery />
        </div>
    </div>
  );
};

export default Right;
