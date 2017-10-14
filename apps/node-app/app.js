const Koa = require('koa');
const app = (module.exports = new Koa());

app.use(async function(ctx) {
  ctx.body = {
    response: 'OK',
    by: process.env.hostname || 'node-app'
  };
});

console.log('Listening on port 80');
if (!module.parent) app.listen(80);
