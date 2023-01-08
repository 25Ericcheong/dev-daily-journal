# Track my Stash
## Project Description
Web app project to track my ~~stash~~ finance haha.

I plan on utilizing the MERN stack and the following frameworks:

Tailwind CSS <br>
Express.js

Still need to connect to Mongo DB with Mongoose.

Can follow this guideline https://blog.logrocket.com/mern-stack-tutorial/ and https://blog.logrocket.com/how-to-set-up-node-typescript-express/

<h2>Documentation</h2>
<h3>NPM</h3>
<h4>Setting up builds</h4>
* Had a bit of a hiccup when trying to setup the backend. Realized that the scripts command be executed were not correctly set. This will be used as a reference to explain why specific commands were used based on attached source (a bit more researching done for my benefit). <br><br/>
* <code>npx tsc</code> compiles the TypeScipript back to vanilla JavaScript. To ensure we do not need to run this command over again, we can ensure this is watched and include the following command <code>npx tsc --watch</code> which compiles our project whenever any changes are made/saved. <br><br/>
* With <code>concurrently \"npx tsc --watch\" \"nodemon -q dist/\"</code>, <code>concurrently</code> is a package installed to execute 2 commands at the same time which will allow us to watch any changes made in any of our TypeScript files and at the same time watch the output JavaScript files with <code>nodemon -q dist/</code>. The node application will restart when any changes are detected and <code>-q</code> is the quiet command that minimises nodemon messages to start/stop only which can be found in nodemon's <code>/doc/cli/options.txt</code><br><br/>
* Source: https://blog.logrocket.com/how-to-set-up-node-typescript-express/
