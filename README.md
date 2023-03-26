# Dev daily journal
## Project Description
Web app project to log what I've learnt and my daily achievements! I haven't really been able to log these achievements in a more structured manner and I believe by having a personalized or internal tool - I would be able to do this more!

## Goals
- Naturally, this web app or tool should enable a user to create a log, categorize the log as an achievement or a simple note (also attach tags to them to help with searching later on) and save these newly created logs in a database. User should be able to delete or modify these logs as well.
- Use and familiarize Node.js as the backend and ensure React can be rendered in the server (with TypeScript as well if possible). In addition to this, I also would like to learn PostgreSQL too (or Mongo in this case).

## Timeline
- [ ] Project configuration and setup [19/03/2023]
- [ ] To be able to create, modify, acquire and delete existing logs [02/04/2023]
- [ ] A starting page that asks user to create an account or allow them to look at my "achievements" [09/04/2023]
- [ ] Deployment and bug fixes if any pops up [23/04/2023]

## Bonus
- [ ] Setting up tools for testing that tests code before pushing things into production. Possibly have an isolated testing environment first [???]
- [ ] Continuous integration to ensure tests are executed and provide feedback if something goes wrong [???]

## Structure
- The Odin project will be the guide that I will follow as I slowly build this project and with progress - I will document each step of it for reference and at the same time; data to be inserted into this project haha. 

<br/><br/>

## Mongo DB introduction
- Materials are obtained from the [MongoDB](https://learn.mongodb.com/learning-paths/introduction-to-mongodb) official website. This section would consist of short pointers I would like to refer back to in the future if needed and will probably be included in this web app too.

<details><summary><h4>Intro to MongoDB</h4></summary>
<ul>
  <li>Notable keywords; CRUD with MongoDB, search experience, aggregation, indexing, data modeling and transactions</li>
  <li>All exercises are done via IDE on MongoDB course (which is connected to an Atlas cluster)</li>
</ul>
</details>

<details><summary><h4>Getting Started with MongoDB Atlas</h4></summary>
    <details>
    <summary><h4>Introduction to MongoDB - Developer Data Platform</h4></summary>
      <ul>
        <li>Database as a serivce (DBaaS) - do not need to configure or manage database but Atlas would do this for us</li>
        <li>Replica Set: data is stored in more than one server (a group of server that holds data). Redundancy and availability</li>
        <li>Type of instances: Clusters (Shared & Dedicated) - serveral mongo servers working  together</li>
        <li>Type of instances: Serverless - charge based on usage and will scale depending  on needs</li>
        <li>Has data API, graphQL API and triggers to enable event driven architecture</li>
      </ul>
    </details>
    <details>
    <summary><h4>Creating and deploying at Atlas cluster</h4></summary>
      <ul>
        <li>Organizations - define and manager userss and teams</li>
        <li>Projects - create separate projects for development, testing and production</li>
        <li>Add database user for SSH purposes</li>
        <li>Remember add own IP address (or any other) so that MongoDB Atlas doesn't block as it blocks everything but anything within Atlas themselves</li>
        <li>Data explorer can be used to visualize existing data in clusters</li>
      </ul>
    </details>
</details>
