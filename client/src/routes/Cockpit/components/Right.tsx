import Odometer from "../../../components/Odometer";
import Battery from "../../../components/Battery";


const Right = () => {
  return (
    <div className="right">
        <div className="right__center">
            <Odometer />
        </div>
        <div className="right__bottom">
          <Battery />
        </div>
    </div>
  );
};

export default Right;
