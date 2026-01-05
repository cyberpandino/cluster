import React, { useRef, useEffect } from 'react';
import { Canvas, useFrame, useThree } from '@react-three/fiber';
import { OrbitControls, useGLTF } from '@react-three/drei';
import { Mesh } from 'three';
import { useSnapshot } from "valtio";
import { state } from "../../store/state";
import { graphics } from "../../config/environment";

function Model() {
    const gltf = useGLTF('/panda.glb');

    return (
        <primitive
            object={gltf.scene}
            dispose={null}
        />
    );
}

function WireframeModel() {
    const gltf = useGLTF('/panda.glb');

    gltf.scene.traverse((child) => {
        if ((child as Mesh).isMesh) {
            const mesh = child as Mesh;
            if (mesh.material) {
                if (Array.isArray(mesh.material)) {
                    mesh.material.forEach((mat: any) => {
                        mat.wireframe = true;
                    });
                } else {
                    (mesh.material as any).wireframe = true;
                }
            }
        }
    });

    return <primitive object={gltf.scene} dispose={null} />;
}

function CameraSetup() {
    const { camera } = useThree();
    const snapSpeed = useSnapshot(state.session.speed);
    const currentSpeed = snapSpeed.current || 0;
    
    // Configurazioni telecamera
    const stationaryConfig = {
        position: [-43.860286822418644, 3.6618990850631516, -99.2060346810957],
        rotation: [-3.1046973443235064, -0.41602413287085493, -3.126676602884795]
    };
    
    const movingConfig = {
        position: [0, 108.5309706129502, 0], // z verrà calcolato dinamicamente
        rotation: [-1.5707953269656727, -1.848033301214691e-8, 0]
    };
    
    useFrame(() => {
        if (currentSpeed === 0) {
            // Animazione verso posizione ferma
            camera.position.lerp(
                { x: stationaryConfig.position[0], y: stationaryConfig.position[1], z: stationaryConfig.position[2] },
                0.05
            );
            camera.rotation.x += (stationaryConfig.rotation[0] - camera.rotation.x) * 0.05;
            camera.rotation.y += (stationaryConfig.rotation[1] - camera.rotation.y) * 0.05;
            camera.rotation.z += (stationaryConfig.rotation[2] - camera.rotation.z) * 0.05;
        } else {
            // Calcola posizione Z in base alla velocità (0-120 km/h → 0-107.025)
            const speedRatio = Math.min(currentSpeed / 120, 1);
            const targetZ = speedRatio * 107.02521009043419;
            
            // Animazione verso posizione in movimento
            camera.position.lerp(
                { x: movingConfig.position[0], y: movingConfig.position[1], z: targetZ },
                0.05
            );
            camera.rotation.x += (movingConfig.rotation[0] - camera.rotation.x) * 0.05;
            camera.rotation.y += (movingConfig.rotation[1] - camera.rotation.y) * 0.05;
            camera.rotation.z += (movingConfig.rotation[2] - camera.rotation.z) * 0.05;
        }
    });

    return null;
}

function CameraLogger() {
    const { camera } = useThree();
    const snapSpeed = useSnapshot(state.session.speed);
    
    useFrame(() => {
        // Logga la posizione della telecamera e velocità (opzionale per debug)
        // console.log('Posizione telecamera:', {
        //     speed: snapSpeed.current,
        //     x: camera.position.x,
        //     y: camera.position.y,
        //     z: camera.position.z,
        //     rotation: {
        //         x: camera.rotation.x,
        //         y: camera.rotation.y,
        //         z: camera.rotation.z
        //     }
        // });
    });

    return null;
}

export default function ModelViewer() {
    // Quality < 3: mostra immagine statica invece del modello 3D
    if (graphics.quality < 3) {
        return (
            <div className="modelViewer-static" style={{
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center',
                width: '100%',
                height: '100%',
            }}>
                <img 
                    src="/fallback-3d-model.png" 
                    alt="Panda 3D" 
                    style={{
                        marginTop: '4%',
                        width: '65%',
                        height: '65%',
                        objectFit: 'contain',
                        opacity: 0.6,
                    }}
                />
            </div>
        );
    }

    return (
        <Canvas 
            gl={{ alpha: true }}
            camera={{ 
                position: [-43.860286822418644, 3.6618990850631516, -99.2060346810957],
                rotation: [-3.1046973443235064, -0.41602413287085493, -3.126676602884795]
            }}>
            <ambientLight />
            <pointLight position={[10, 10, 10]} />
            <WireframeModel />
            <OrbitControls 
                enablePan={false}
            />
            <CameraSetup />
            <CameraLogger />
        </Canvas>
    );
}
