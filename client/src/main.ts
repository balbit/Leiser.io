let canvas = <HTMLCanvasElement> document.getElementById("myCanvas")!;
let ctx = canvas.getContext("2d")!;

let upPressed: boolean = false;
let downPressed: boolean = false;
let leftPressed: boolean = false;
let rightPressed: boolean = false;

document.addEventListener("keydown", (ev: KeyboardEvent) => {
    switch (ev.key) {
        case "w":
            upPressed = true;
            break;
        case "a":
            rightPressed = true;
            break;
        case "s":
            downPressed = true;
            break;
        case "d":
            leftPressed = true;
            break;
    }
})

document.addEventListener("keyup", (ev: KeyboardEvent) => {
    switch (ev.key) {
        case "w":
            upPressed = false;
            break;
        case "a":
            rightPressed = false;
            break;
        case "s":
            downPressed = false;
            break;
        case "d":
            leftPressed = false;
            break;
    }
})

class Player {
    public x: number;
    public y: number;

    constructor(startX: number, startY: number) {
        this.x = startX;
        this.y = startY;

        console.log(this.x, this.y);
    }

    render () {
        console.log(`${this.x}, ${this.y}`);
        ctx.clearRect(0, 0, 600, 400);
        ctx.fillStyle = "#d3d3d3";
        ctx.beginPath();
        ctx.arc(this.x, this.y, 20, 0, 2 * Math.PI);
        ctx.fill();
    }
}

let ws : WebSocket;
let player : Player;

function main() {

    console.log("Client starting!");

    ws = new WebSocket("ws://localhost:3000/ws");
    
    ws.onopen = function() {
        console.log("WebSocket connection established");
        ws.send("init");
    };

    ws.onerror = function(error) {
        console.log("WebSocket Error: ", error);
    };

    ws.onmessage = function(ev: MessageEvent) {
        let msg: string[] = ev.data.split(' ');
        console.log(msg);
        switch (msg[0]) {
            case "start_pos":
                player = new Player(parseInt(msg[1]), parseInt(msg[2]));
        }
    }
}

function loop () {
    player.render();
    let msg: string = `keys ${+ upPressed} ${+ downPressed} ${+ leftPressed} ${+ rightPressed}`
    if (upPressed || downPressed || leftPressed || rightPressed) ws.send(msg)
}


document.addEventListener("DOMContentLoaded", main);
window.setInterval(loop, 60)