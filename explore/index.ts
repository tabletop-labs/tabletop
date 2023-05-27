import Bao from "baojs";

const SERVER_PORT = parseInt(process.env.SERVER_PORT || "3000", 10);
const SERVER_HOSTNAME = process.env.SERVER_HOSTNAME || "localhost";

const app = new Bao();

app.get("/", (ctx) => {
  return ctx.sendText("Welcome to explore Bao.js!");
});

const server = app.listen({
  hostname: SERVER_HOSTNAME,
  port: SERVER_PORT
});

console.log(`Listening on ${server.hostname}:${server.port}`);
