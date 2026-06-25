import "./WarningLights.scss";
import { useSnapshot } from "valtio";
import { state } from "../../store/state";
import { useEffect, useRef, useState } from "react";

const WarningLights = () => {
  const snapWarnings = useSnapshot(state.warnings);
  const [lastShownIcon, setLastShownIcon] = useState<string | null>(null);
  const prevWarningsRef = useRef<typeof state.warnings>({ ...state.warnings });
  const hideTimeoutRef = useRef<ReturnType<typeof setTimeout> | null>(null);

  const warningColors = {
    doors: '#FFA500',
    light: '#FFFF00',
    lowBeam: '#00FF00',
    highBeam: '#0000FF',
    fogLight: '#00FF00',
    engineCoolant: '#FF0000',
    warning: '#FFA500',
    hazard: '#FFA500',
    turnSignals: '#00FF00',
    battery: '#FF0000',
    brakeSystem: '#FF0000',
    fuel: '#FFFF00',
    injectors: '#FF0000',
    keyOn: '#0000FF',
    rearDefrost: '#00FF00',
    engineOil: '#FF0000',
  };

  useEffect(() => {
    for (const key of Object.keys(snapWarnings)) {
      const k = key as keyof typeof snapWarnings;
      if (snapWarnings[k] && !prevWarningsRef.current[k]) {
        setLastShownIcon(k);
        if (hideTimeoutRef.current) clearTimeout(hideTimeoutRef.current);
        hideTimeoutRef.current = setTimeout(() => setLastShownIcon(null), 5000);
        break;
      }
    }
    prevWarningsRef.current = { ...snapWarnings };
  }, [snapWarnings]);

  useEffect(() => {
    return () => {
      if (hideTimeoutRef.current) clearTimeout(hideTimeoutRef.current);
    };
  }, []);

  return (
    <div className="componentWarningLights">
      <ul>
        {Object.entries(snapWarnings).map(([key, value]) => (
          <li key={key}>
            <img
              src={`/icons/warnings/${key}.svg`}
              className={value ? 'active' : ''}
              alt={key}
              style={value ? {
                filter: `drop-shadow(0 0 8px ${warningColors[key as keyof typeof warningColors]})`,
                opacity: 1
              } : undefined}
            />
          </li>
        ))}
      </ul>

      {lastShownIcon && (
        <div className="warningOverlay">
          <div className="warningOverlay__wrapper">
            <img
              src={`/icons/warnings/${lastShownIcon}.svg`}
              alt={lastShownIcon}
              style={{
                filter: `drop-shadow(0 0 12px ${warningColors[lastShownIcon as keyof typeof warningColors]})`,
              }}
            />
          </div>
        </div>
      )}
    </div>
  );
};

export default WarningLights;
