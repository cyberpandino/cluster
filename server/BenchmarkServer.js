const http = require('http');
const socketIo = require('socket.io');

const port = 3001;

const server = http.createServer((req, res) => {
    res.writeHead(200);
    res.end('Panda Benchmark Server Active');
});

const io = socketIo(server, {
    cors: { origin: "*", methods: ["GET", "POST"] }
});

// Variabili simulazione
let rpm = 800;
let speed = 0;
let direction = 1;

io.on('connection', (socket) => {
    console.log(`CLIENT RICONOSCIUTO: ${socket.id}`);
    console.log("\n Inizio invio dati ...");

    const timer = setInterval(() => {
        // Logica animazione
        rpm += (80 * direction);
        speed += (0.8 * direction);

        if (rpm > 5500) direction = -1;
        if (rpm < 800) direction = 1;

        const complexPacket = {
            parameters: {
                rpm: { 
                    value: Math.round(rpm), 
                    success: true 
                },
                speed: { 
                    value: Math.floor(Math.abs(speed)), 
                    success: true 
                },
                coolant_temp: { 
                    value: 90, 
                    success: true 
                },
                engine_load: {
                    value: 25,
                    success: true
                },
                battery_voltage: {
                    value: 13.8,
                    success: true
                }
            }
        };

        // Invio su obd-live
        socket.emit('obd-live', complexPacket);
        
        // Debug
        if (rpm > 5400 || rpm < 850) {
             process.stdout.write(".");
        }

    }, 50); // 20 volte al secondo

    socket.on('disconnect', () => {
        console.log(`\n Client disconnesso`);
        clearInterval(timer);
    });
});

server.listen(port, () => {
    console.log(` BENCHMARK SERVER (FIXED JSON) SU PORTA ${port}`);
});
