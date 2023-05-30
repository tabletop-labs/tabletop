import Bao from "baojs";

const SERVER_PORT = parseInt(process.env.SERVER_PORT || "3000", 10);
const SERVER_HOSTNAME = process.env.SERVER_HOSTNAME || "localhost";

const app = new Bao();

app.get("/", (ctx) => {
  return ctx.sendText("Welcome to console Bao.js!");
});

const router = new Bun.FileSystemRouter({
  style: "nextjs",
  dir: "./pages",
  origin: "https://mydomain.com",
  assetPrefix: "_next/static/"
});
router.match("/");
app.use(router);

app.errorHandler = (_error) => {
  return new Response("Oh no! An error has occurred...");
};

app.notFoundHandler = (_ctx) => {
  return new Response("Route not found...");
};

const server = app.listen({
  hostname: SERVER_HOSTNAME,
  port: SERVER_PORT
});

console.log(`Listening on ${server.hostname}:${server.port}`);
