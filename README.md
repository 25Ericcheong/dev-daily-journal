# Dev daily journal
## Project Description
Web app project to log what I've learnt and my daily achievements! I haven't really been able to log these achievements in a more structured manner and I believe by having a personalized or internal tool - I would be able to do this more!

## Goals
- Naturally, this web app or tool should enable a user to create a log, categorize the log as an achievement or a simple note (also attach tags to them to help with searching later on) and save newly created logs in a database. User should be able to delete or modify logs as well.
- Use and familiarize Node.js as the backend and ensure React can be rendered in the server (with TypeScript as well if possible). In addition to this, I also would like to learn PostgreSQL too (or Mongo in this case).

## Timeline
- [ ] Project configuration and setup [19/03/2023]
- [ ] To be able to create, modify, acquire and delete existing logs [02/04/2023]
- [ ] A starting page that asks user to create an account or allow them to look at my "achievements" [09/04/2023]
- [ ] Deployment and bug fixes if any pops up [23/04/2023]

## Bonus
- [ ] Setting up tools for testing that tests code before pushing things into production. Possibly have an isolated testing environment first [???]
- [ ] Continuous integration to ensure tests are executed and provide feedback if something goes wrong [???]
- [ ] Completing MongoDB introduction course on MongoDB university

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
  <ol>
    <li><details><summary><h4>Introduction to MongoDB - Developer Data Platform</h4></summary>
      <ul>
        <li>Database as a serivce (DBaaS) - do not need to configure or manage database but Atlas would do this for us</li>
        <li>Replica Set: data is stored in more than one server (a group of server that holds data). Redundancy and availability</li>
        <li>Type of instances: Clusters (Shared & Dedicated) - serveral mongo servers working  together</li>
        <li>Type of instances: Serverless - charge based on usage and will scale depending  on needs</li>
        <li>Has data API, graphQL API and triggers to enable event driven architecture</li>
      </ul>
    </details></li>
    <li><details>
    <summary><h4>Creating and deploying at Atlas cluster</h4></summary>
      <ul>
        <li>Organizations - define and manager userss and teams</li>
        <li>Projects - create separate projects for development, testing and production</li>
        <li>Add database user for SSH purposes</li>
        <li>Remember add own IP address (or any other) so that MongoDB Atlas doesn't block as it blocks everything but anything within Atlas themselves</li>
        <li>Data explorer can be used to visualize existing data in clusters</li>
      </ul>
    </details></li>
  </ol>
</details>

<details><summary><h4>MongoDB and the Document Model</h4></summary>
  <ol>
    <li><details><summary><h4>MongoDB database</h4></summary>
      <ul>
        <li>MongoDB stored as documents similar to JSON - able to use one format for any applicaitons</li>
        <li>Document - basic unit of data in MongoDB</li>
        <li>Collection - set of documents (structure may not be the same within collection since a document has flexible schema)</li>
        <li>Database - set of collections</li>
        <li>Atlas offers full text search and data visualization</li>
     </ul>
    </details></li>
    <li><details><summary><h4>MongoDB Document Model</h4></summary>
      <ul>
        <li>Documents displayed in JSON and stored in BSON (extension of JSON but in binary which supports multiple different data types) on database.</li>
        <li>Every document requires <code>_id</code> field, if document doesn't have one - MongoDB auto generates ObjectId to represent <code>_id</code> field</li>
        <li>Collections can have different documents schemas in it (because schema is flexbile as it supports polymorphic documents).</li>
        <li>To have more control over database, can have optional schema validation in placed. This can be used to constraint structure of documents.</li> 
       </ul>
    </details></li>
    <li><details><summary><h4>Managing Databases, Collections, and Documents in Atlas Data Explorer</h4></summary>
      <ul>
        <li>Atlas Data Explorer - can create collections or databases and insert documents into databases as required</li>
        <li>Atlas UI - useful for testing purposes</li>
       </ul>
    </details></li>
  </ol>
</details>

<details><summary><h4>Connecting to a MongoDB Database</h4></summary>
  <ol>
    <li><details><summary><h4>Using MongoDB connection strings</h4></summary>
      <ul>
        <li>Connection strings allows us to connect to cluster and work with data. It describes host that we will be using and the options for connecting to MongoDB database</li>
        <li>Connecting string can be used to connect from mongo shell, mongo compass or to any other app</li>
        <li>There exists two formats of connecting string - standard format and DNS seed list format</li>
        <li>Standard format used to connect to standalone clusters, replica sets or sharded clusters</li>
        <li>DNS seed list format provides a DNS server list to connection string. This provides flexibility of deployment and can change servers in rotation without reconfiguring clients.</li>
        <li>Connection string consists of username and password (created database users that have access to database), host and optional port number to database and lastly, additional options</li>
     </ul>
    </details></li>
    <li><details><summary><h4>Connecting to a MongoDB Atlas Cluster with the Shell</h4></summary>
      <ul>
        <li>Will require mongosh to connect into our cluster locally via CLI with provided connection string</li>
        <li>mongosh is a Node.js REPL environment and will enable us to utilize JavaScript languages within the shell</li>
     </ul>
    </details></li>
    <li><details><summary><h4>Connecting to a MongoDB Atlas Cluster with MongoDB Compass</h4></summary>
      <ul>
        <li>GUI to allow us to query and analyze data in cluster</li>
        <li>Enable us to acquire statistical summary of databases existing in cluster</li>
     </ul>
    </details></li>
    <li><details><summary><h4>Connecting to a MongoDB Atlas Cluster from an Application</h4></summary>
      <ul>
        <li>MongoDB drivers allow us to connect database to application using programming language of our choice with provided connection string</li>
        <li>MongoDB documentation to use for configuration and setting up connection to MongoDB database</li>
     </ul>
    </details></li>
    <li><details><summary><h4>Troubleshooting MongoDB Atlas connection errors</h4></summary>
      <ul>
        <li>Network access errors - can be due to not adding IP address for network access</li>
        <li>User authentication errors - can be due to not including password</li>
     </ul>
    </details></li>
  </ol>
</details>

<details><summary><h4>MongoDB CRUD Operations: Insert and Find Documents</h4></summary>
  <ol>
    <li><details><summary><h4>Inserting documents in a MongoDB Collection</h4></summary>
      <ul>
        <li><code>insertOne()</code> and <code>insertMany()</code> are the relevant code to insert documents.</li>
        <li>Note that with the use of <code>insertOne()</code>, if the collection targeted for document insertion does not exist. It will create the collection automatically. This is worth noting as it causes collections to be created unnecessarily.</li>
     </ul>
    </details></li>
    <li><details><summary><h4>Finding documents in a MongoDB Collection</h4></summary>
      <ul>
        <li><code>use database-name</code> will allow us to utilize the database has the relevant collections included. Next, <code>db.collection-name.find()</code> will return documents that can be found in the collection-name inputted.</li>
        <li>To specify what fields/values a document should have - <code>db.collection-name.find({ field-name: value })</code> can be used. This ensures that we want a specific key/field name to contain a specific value</li>
        <li>To target multiple value for a key/field name - <code>db.collection-name.find({ field-name: { $in: [value-one, value-two] } })</code> can be considered where the values we are looking for can be value-one or value-two.</li>
      </ul>
    </details></li>
    <li><details><summary><h4>Finding documents in a MongoDB by using Comparison Operators</h4></summary>
      <ul>
        <li>Examples of comparison operators can be as such <code>$gt</code>, <code>$lt</code>, <code>$gte</code> and <code>$lte</code></li>
        <li>An example of utilizing a comparison operator is <code>db.collection-name.find({ field-name: { $gt: 50 } })</code></li>. Note that in this case <code>field-name</code> could just be a field-name specifically or object.field-name (object is the sub document within a document) - depending on the structure of the document.
        <li>Note that if there is an array of items and even if one item fulfills the comparison operator or condition - the document is retrieved along with the entire array of items (even ones that do not fulfill the condition)</li>
        <li>Those are not the only comparison operator that exists.</li>
      </ul>
    </details></li>
    <li><details><summary><h4>Querying on Array Elements in MongoDB</h4></summary>
      <ul>
        <li>Notice the difference between the two queries. <code>db.collection-name.find({field-name: value})</code> and <code>db.collection-name.find({field-name: {$elemMatch: {$eq: value}}})</code>. The first will return any item with the field that has the scalar value (value does not need to be in an array). The second will return a items with the matching values found in an array for the specified field. The difference is that the second will require values to be in an array.</li>
        <li>Can use <code>$elemMatch</code> to also acquire items that meet the query criteria included. Example, <code>db.collection-name.findOne({sub-document-name: {$elemMatch: {field-name: value, another-field-name: {$lt : value}, other-field-name: {$gte : another value}}}})</code>. This example will provided us with one item that meets the three queries/conditions listed for three separate fields.</li>
      </ul>
    </details></li>
    <li><details><summary><h4>Finding Documents by Using Logical Operators</h4></summary>
      <ul>
        <li>The two logical operators will be $and & $or operators. Example, <code>db.collection-name.find({$and: [{field-name: value}, {another-field-name: {$gte: greater-value}}]})</code>. This can also be accomplished implicitly via <code>db.collection-name.find({field-name: value, another-field-name: {$gte: greater-value}})</code>.</li>
        <li>The $or operator is exactly the same syntax. You can also combine a single query with multiple $or and $and operator together. Note that explicit $and is required first when intention is to have two same $or being used together as the first $or operator will override the subsequent operator as it defies the JSON object structure with all keys required to be unique.</li>
        <li>Example for better visualization 
          <pre><code>
          db.routes.find({
            $and: [
                { $or: [{ dst_airport: "SEA" }, { src_airport: "SEA" }] },
                { $or: [{ "airline.name": "American Airlines" }, { airplane: 320 }] },
              ]
            })
         </code></pre>
        </li>
        <li>Example from assignment to access a sub document with a comparison operator within it - <code>db.sales.find({ couponUsed: true,  purchaseMethod: "Online", "customer.age": { $lte: 25 } })</code></li>
        <li>
          <b>Note.</b> If there are sub documents present, and you are looking to check that the sub-documents has a specific value with the specified field - use the dot notation method to do this. There is a big difference between both code blocks included below
          <pre><code>
          db.sales.find({
            $or: [{ items: {name: "notepad", tags: "school"}}]
          })
          </pre></code>
          <pre><code>
          db.sales.find({
            $or: [{ "items.name": "pens" }, { "items.tags": "writing" }],
          })
          </code></pre>
          The first is we have an <code>$or</code> comparison operator (which in this case isn't really neccessary anymore) to look for an items sub document to have that specific object (if any other items sub documents with object structure that differs from this - it will not meet the query's expression then and will not show up on results). The second in this case looks for <b>any</b> items sub document with either of the fields (checks if it exists) along with the value provided. If any is true, the document will show up regardless of the <a href="https://stackoverflow.com/questions/38129635/mongodb-accessing-subdocuments" target="_blank">object structure</a>.
        </li>
      </ul>
    </details></li>
  </ol>
</details>

<details><summary><h4>MongoDB CRUD Operations: Replace and Delete Documents</h4></summary>
  <ol>
    <li><details><summary><h4>Replacing a Document in MongoDB</h4></summary>
      <ul>
        <li>Using the <code>db.collection-name.replaceOne(filter, replacement, options)</code> will allow us to replace a document within a collection.</li>
        <li>In this case, <code>filter</code> would typically consist of a way to uniquely identify a document within the collection. Additionally, <code>replacement</code> would contain an updated version of the document (excluding its id field) to replace the existing document within our collection.</li>
      </ul>
    </details></li>
    <li><details><summary><h4>Updating a Document in MongoDB</h4></summary>
      <ul>
        <li>Commands to use - <code>updateOne(filter, update, options)</code> along with update operators - <code>$set</code> and <code>$push</code> along with <code>upsert</code> will be taught in this section which are used within the parameters of the <code>updateOne</code> method provided</li>
        <li><code>$set</code> operator can be used to add new fields or values to a document or replace existing fields with new values in a document</li>
        <li><code>$push</code> has the same capabilities as well and it also appends new item to an existing array value or if the array value does not exist in the first place, it creates an array with the new item to be added</li>
        <li>If the filter query filled does not provide any existing document within collection, we would use <code>upsert</code> to create a new document if that were to be the case. Upsert stands for update or insert and is included as an option object which is set to be true or false</li>
        <li>Some examples for better utilization of methods
          Look for specified id and set subscribers field to the value provided
          <pre><code>
          db.podcasts.updateOne(
            {_id: ObjectId("5e8f8f8f8f8f8f8f8f8f8f8")},
            {$set: {subscribers: 98562}}
          )
          </pre></code>
          Look for a document with the specified title value and set the topics value with the provided array. If it does not exist, insert a new document (purpose of upsert). You can also use <code>$inc</code> followed by the field and the value to incrementally increase the existing value by as well if required
          <pre><code>
          db.podcasts.updateOne(
            { title: "The Developer Hub" },
            { $set: { topics: ["databases", "MongoDB"] } },
            { upsert: true }
          )
          </pre></code>
          Add a new item to existing value array of hosts field
          <pre><code>
          db.podcasts.updateOne(
            { _id: ObjectId("5e8f8f8f8f8f8f8f8f8f8f8") },
            { $push: { hosts: "Nic Raboy" } }
          )
          </pre></code>
          To add multiple items to an array (if item does not exist - nothing will happen), you would require the <code>$each</code> command as shown below.
          <pre><code>
          db.birds.updateOne(
            { _id: ObjectId("6268471e613e55b82d7065d7") },
            {
              $push: {
                diet: { $each: ["newts", "opossum", "skunks", "squirrels"] },
              },
            }
          )
          </pre></code>
        </li>
      </ul>
    </details></li>
    <li><details><summary><h4>Updating a few Documents in MongoDB</h4></summary>
      <ul>
        <li></li>
      </ul>
    </details></li>
  </ol>
</details>
