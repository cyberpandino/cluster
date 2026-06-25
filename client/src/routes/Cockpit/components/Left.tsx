import Tachometer from "../../../components/Tachometer";
import Altitude from "../../../components/Altitude";


const Left = () => {
  return (
    <div className="left">
        <div className="left__center">
            <Tachometer  />
        </div>
        <div className="left__bottom">
          <Altitude />
        </div>
    </div>
  );
};

export default Left;
