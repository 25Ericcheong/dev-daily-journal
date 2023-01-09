# Track my Stash
## Project Description
Web app project to track my ~~stash~~ finance haha.

I plan on utilizing the MERN stack and the following frameworks:
* Tailwind CSS <br>
* Express.js

Still need to connect to Mongo DB with Mongoose and ensure that a connection is made from the client (frontend) to the server (backend).

Can follow this guideline https://blog.logrocket.com/mern-stack-tutorial/

## Documentation
### NPM Related
<details>
<summary>#### Setting up for Development</summary>
This is a dropdown with text!
</details>
#### Setting up for Development
* Had a bit of a hiccup when trying to setup the backend. Realized that the scripts command be executed were not correctly set. This will be used as a reference to explain why specific commands were used based on attached source (a bit more researching done for my benefit). 
* `npx tsc` compiles the TypeScipript back to vanilla JavaScript. To ensure we do not need to run this command over again, we can ensure this is watched and include the following command `npx tsc --watch` which compiles our project whenever any changes are made/saved. 
* With `concurrently \"npx tsc --watch\" \"nodemon -q dist/\"`, `concurrently` is a package installed to execute 2 commands at the same time which will allow us to watch any changes made in any of our TypeScript files and at the same time watch the output JavaScript files with `nodemon -q dist/`. The node application will restart when any changes are detected and `-q` is the quiet command that minimises nodemon messages to start/stop only which can be found in nodemon's `/doc/cli/options.txt`
* Source: https://blog.logrocket.com/how-to-set-up-node-typescript-express/
