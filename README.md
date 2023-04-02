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
    <li><details><summary><h4>Updating a Document and acquiring updated version in MongoDB</h4></summary>
      <ul>
        <li>There is a difference between <code>findAndModify({query: {filter-object}, update: {updated-fields-object}, new: true})</code> and <code>updateOne()</code> methods. The first method is used to return document that has just been updated. The <code>new: true</code> ensures that we get an updated version of the document as a return with the first method.</li>
        <li>Typically, we would use <code>updateOne()</code> and <code>findOne()</code> to update a document and then look for the updated document. This does two round trips to and back from the server. This could also return the same document that could have already been outdated to begin with before returning a version that you thought is the most updated version.</li>
        <li>An example is as shown</li>
        <pre><code>
        db.podcasts.findAndModify({
          query: { _id: ObjectId("6261a92dfee1ff300dc80bf1") },
          update: { $inc: { subscribers: 1 } },
          new: true,})
        </pre></code>
        Note that you can also set <code>upsert: true</code> to insert a new document if no documents are acquired with the inputted filter query.
      </ul>
    </details></li>
    <li><details><summary><h4>Updating multiple documents in MongoDB</h4></summary>
      <ul>
        <li>Using <code>updatrMany({filter-object}, {updated-fields-object})</code> will allow the updates to occur for all documents that meet the filter criteria in the filter object</li>
        <li>This method does not guarantee all documents will be updated at the same time. This can be checked against the number of successful update counts against the matched count documents.</li>
        <li>Example is as shown. This code will update all documents with any of the values found in the array provided with the use of <code>$in</code> and then update any documents found with the specified date for the given field name.</li>
        <pre><code>
        db.birds.updateMany(
          {
            common_name: { $in: ["Blue Jay", "Grackle"]},
          },
          {
            $set: {last_seen: ISODate("2022-01-01")},
          }
        )
        </pre></code>
      </ul>
    </details></li>
    <li><details><summary><h4>Removing one or many documents in MongoDB</h4></summary>
      <ul>
        <li>You can choose to delete a document with <code>deleteOne({filter-object})</code> or many documents with <code>deleteMany({filter-object})</code></li>
        <li>Example is as shown to delete one document</li>
        <pre><code>
        db.birds.deleteOne({_id: ObjectId("35465")})
        </pre></code>
        <li>Example is as shown to delete multiple documents</li>
        <pre><code>
        db.birds.deleteMany({ sightings_count: { $lte: 10 } })
        </pre></code>
      </ul>
    </details></li>
  </ol>
</details>

<details><summary><h4>MongoDB CRUD Operations: Modifying Query Results</h4></summary>
  <ol>
    <li><details><summary><h4>Sorting and Limiting Query Results in MongoDB</h4></summary>
      <ul>
        <li><code>cursor</code> is a pointer to the result set of a query like <code>find</code>. The methods that comes with it would be <code>.sort({field-name: 1 or -1})</code> and <code>.limit(size-of-results)</code>. Note that these are chained on top of a query. In the sort method, 1 would sort results from smallest to largest alphabetically (Mongo sorts by capitalized letters first)</li>
        <li>Limiting number of results can improve performance of app.</li>
        <li>Example of utilizing limit and sort methods. To sort by most recent dates - use <code>-1</code> instead</li>
        <pre><code>
        db.companies
          .find({ category_code: "music" })
          .sort({ number_of_employees: -1})
          .limit(3)
        </pre></code>
        <li>An additional interesting example here as well which looks for any of the included names in provided array</li>
        <pre><code>
        db.sales
        .find({ "items.name": { $in: ["laptop", "backpack", "printer paper"] }, "storeLocation": "London", })
        .sort({ saleDate: -1, })
        .limit(3)
        </pre></code>
      </ul>
    </details></li>
    <li><details><summary><h4>Returning specific data from a query in MongoDB</h4></summary>
      <ul>
        <li>To limit specific fields from being acquired which will improve bandwidth performance. This is called projection.</li>
        <li>Can choose to include or exlucde fields (can't combine both exception for _id field). To include we have field values to be 1 while to exclude, we have other field names to be 0 instead.</li>
        <pre><code>
        // Return all restaurant inspections - business name and result fields only
        db.inspections.find(
          { sector: "Restaurant - 818" },
          { business_name: 1, result: 1, _id: 0 }
        )
        </pre></code>
        and even this - since we are excluding sub document fields
        <pre><code>
        // Return all inspections with result of "Pass" or "Warning" - exclude date and zip code
        db.inspections.find(
          { result: { $in: ["Pass", "Warning"] } },
          { date: 0, "address.zip": 0 }
        )
        </pre></code>
      </ul>
    </details></li>
    <li><details><summary><h4>Counting documents in a MongoDB Collection</h4></summary>
      <ul>
        <li>This can be done with the <code>countDocuments({query-to-count-specific-documents-object}, options)</code> method. </li>
        <pre><code>
        // Count number of trips over 120 minutes by subscribers
        db.trips.countDocuments({ tripduration: { $gt: 120 }, usertype: "Subscriber" })
        </pre></code>
      </ul>
    </details></li>
  </ol>
</details>

<details><summary><h4>MongoDB Aggregation</h4></summary>
  <ol>
    <li><details><summary><h4>Inroduction to MongoDB Aggregation</h4></summary>
      <ul>
        <li>Used to build multi stage query - a series of stages completed one at a time, in order. Within each stage, data can be filtered, sorted, grouped and transformed.</li>
        <li>The following can be used <code>.aggegate([{ $stage_name: {expression} }, { $second_stage_name: {expression} }])</code></li>
        <li>Stage is a single operation on data. Commonly used in stages - <code>$match</code> (for filtering), <code>$group</code> (group documents based on criteria) and <code>$sort</code> (put documents in specified order)</li>
        <li>Field references can be used to acquire and combine values from existing fields in the documents to create new ones if required</li>
        <li>Example as shown below</li>
        <pre><code>
        db.collection.aggregate([
            {
                $stage1: {
                    { expression1 },
                    { expression2 }...
                },
                $stage2: {
                    { expression1 }...
                }
            }
        ])
        </pre></code>
      </ul>
    </details></li>
    <li><details><summary><h4>Using $match and $group Stages in a MongoDB Aggregation Pipeline</h4></summary>
      <ul>
        <li>Most commonly used - <code>$match</code> and <code>$group</code>. Best to pplace <code>$match</code> as early as possibly in pipeline so that it can use indexes which helps with processing.</li>
        <li>An example of using <code>$group</code> would be <code>{ $group: {_id: expression, field-name: {accumulator: expression}} }</code></li>
        <pre><code>
        db.zips.aggregate([
          { $match: {state: "CA"} },
          { $group: {_id: "$city", totalZips: { $count : { }}} }
        ])
        </pre></code>
        <li>Additional example for reference. In this case, location is a sub document and coordinates is being used as a grouping key and  we count the number of sightings to acquire famous coordinates</li>
        <pre><code>
        db.sightings.aggregate([
          { $match: {species_common: 'Eastern Bluebird'} }, 
          { $group: {_id: '$location.coordinates', number_of_sightings: 
            { $count: {}} }
          }
        ])
        </pre></code>
      </ul>
    </details></li>
    <li><details><summary><h4>Using $sort and $limit Stages in a MongoDB Aggregation Pipeline</h4></summary>
      <ul>
        <li>Note that order of stages would change the results of aggregation.</li>
        <li>Example for reference</li>
        <pre><code>
        db.zips.aggregate([
          { $sort: {pop: -1} },
          { $limit:  5 }
        ])
        </pre></code>
      </ul>
    </details></li>
    <li><details><summary><h4>Using $project, $count and $set Stages in a MongoDB Aggregation Pipeline</h4></summary>
      <ul>
        <li><code>$project</code> determines output shape and usually last stage since we are formatting output. Can use to include and project fields with 1 or 0. We can even project a new field name too if required.</li>
        <li><code>$set</code> adds or modifies fields in the pipeline. Useful when we want to change existing fields or add new ones to be used in upcoming pipeline stages.</li> 
        <li><code>$count</code> returns total document count. </li>
        <li>Example for <code>$project</code> where a new field name - population is projected with an existing field value. If field name is not included, it would not be projected in results with the exception of _id field</li>
        <pre><code>
        { $project: {state:1, zip:1, population:"$pop", _id:0} }
        </pre></code>
        <li>Example for <code>$set</code> where a new field named place is created and subbed with values coming from two different fields in the existing document</li>
        <pre><code>
        { $set: {
          place: {$concat:["$city",",","$state"]},
          pop: 10000
          }
        }
        </pre></code>
        <li>Example for <code>$count</code> where we count the number of documents in collection and name the output field total_zips</li>
        <pre><code>
        {
          $count: "total_zips"
        }
        </pre></code>
      </ul>
    </details></li>
    <li><details><summary><h4>Using $out Stage in a MongoDB Aggregation Pipeline</h4></summary>
      <ul>
        <li><code>$out</code> has to be the last stage of the pipeline. Writes documents that are returned by an aggregation pipline into a collection. Will create a new collection if it does not exist.</li>
        <li><code>$out: {db: "db-name", coll: "collection-name"}</code>. Can also exclude db field, meaning document is created in the same database that data is being aggregated from currently.</li>
        <li>Note - if collection exists, the database will be overwritten with aggregated results.</li>
      </ul>
    </details></li>
  </ol>
</details>

<details><summary><h4>MongoDB Indexes</h4></summary>
  <ol>
    <li><details><summary><h4>Using MongoDB Indexes in Collections</h4></summary>
      <ul>
        <li>Stores small portion of data from the collection to help with search efficiency. It points to document identity so that it is faster to look up and update specified document</li>
        <li>It speeds up query searches, reduce disk I/O, reduce resources required to execute queries and also supports equality matches as well as range-based operations</li>
        <li>It helps MongoDB such that it wouldn't need to scan entire collection to check that it matches query criteria and preventing the need to sort results in memory</li>
        <li>Every collection will have an index by default (only has <code>_id</code> field). Each query should have its own index.</li>
        <li>Will need to update index data structure of document changes. Ensure that we only have indexes we need and remove unnucessary ones.</li>
        <li>Most commonly used - compound field and single field indexes. Index that perform on arrays are multi key indexes.</li>
      </ul>
    </details></li>
    <li><details><summary><h4>Creating Single Field Index in MongoDB</h4></summary>
      <ul>
        <li>Done with the <code>createIndex()</code> command. Index can also include uniqueness command to prevent insertion of documents with the specified field that should only contain unique values</li>
        <li>There is a way to determine if there are any indexes with the <code>getIndexes()</code> command and can also determine if a query has an index in placed with the <code>explain()</code> command</li>
        <li>Ensuring single field index specified has unique values for each document. Worth noting that if a query is over indexed, it can cause performance issue. Creating single field index with equality constraint</li>
        <li>Note that if index is not being used, you would see a <code>COLLSCAN</code> which indicates that MongoDB had to do a wide collection scan since no index has been created. An <code>IXSCAN</code> indicates an index is being used for specific query.</li>
        <pre><code>
          db.customers.createIndex({ email: 1 },{ unique:true })
        </pre></code>
        <li>To check that index created is working</li>
        <pre><code>
          db.accounts.explain().find(/* your query here */)
        </pre></code>
      </ul>
    </details></li>
    <li><details><summary><h4>Creating Multikey Index in MongoDB</h4></summary>
      <ul>
        <li>This applies to defining an index on an array field which is called a multikey index can index primitive values, sub documents or sub arraays of an array. It can also be a part of a compound field defined index.</li>
        <li>Limitation is that we can only have one array field per index.</li>
      </ul>
    </details></li>
    <li><details><summary><h4>Creating Compound Index in MongoDB</h4></summary>
      <ul>
        <li>Index on multiple fields. Note that if other queries contain even one of the field listed in the multi field index, it will utilize the multi field index created whether it may be prefix or not in a chain of queries.</li>
        <li>Note order matters and it is recommended that the order is as such - equality, range and then sort. This reduces in memory processing time. Meaning, the index created should be created in the order that matches query criteria. This would mean that the placement of fields is important in the index and it needs to be in sync with queries that will utilize the index that will be created.</li>
        <li>An example of ensuring that the order of field for the index created has to match the query criteria as shown below.</li>
        <pre><code>
        db.customers
          .find({ birthdate: {$gte:ISODate("1977-01-01")}, active:true })
          .sort({ birthdate:-1, name:1 })
        </pre></code>
        <li>This ensures that the active field exists first and is true and then sorts birthdate and followed by name</li>
        <pre<code>
        db.customers.createIndex({
          active:1, 
          birthdate:-1,
          name:1
        })
        <li>Ensuring that the query is fully covered ensures that no data is required to be fetched from in-memory or collection. This can be achieved by ensuring that only required fields are projected (as listed in the created index)</li>
        <pre><code>
        db.customers.explain()
        .find(
          { birthdate: {$gte:ISODate("1977-01-01")}, active:true },
          { name:1, birthdate:1, _id:0 })
        .sort({ birthdate:-1, name:1 })
        </pre></code>
      </ul>
    </details></li>
    <li><details><summary><h4>Deleting Indexes in MongoDB</h4></summary>
      <ul>
        <li>Deleting indexes that are no longer required will prevent redundant indexes from being used by a query. This will cause performance issues if not removed properly</li>
        <li>The use of <code>dropIndex()</code> and <code>dropIndexes()</code> will allow the removal of indexes. Note that if no index names are provided to the <code>dropIndexes()</code> method, all indexes will be dropped. Providing an array of index names will remove all indexes provided within that array</li>
      </ul>
    </details></li>
  </ol>
</details>
