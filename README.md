# Dev daily journal
## Project Description
I want to be able to log my daily achievements somewhere. It doesn't need to be achievements but it has to allow me to save what I have learned and allow me to tag these logs so that I could make a wide search for reference in the future. Lastly, a cool feature that I would really want to is to allow people to highlight parts of my article - whether to tell me I have made a grammartical error or an improvement to how I have explained something. 

## Goals
- Naturally, this web app or tool should enable a user to create a log, categorize the log as an achievement or a simple note (also attach tags to them to help with searching later on) and save newly created logs in a database. User should be able to delete or modify logs as well.
- Use and familiarize Golang for the entire project.

## Development (to be moved into project)
- [CompileDaemon](https://github.com/githubnemo/CompileDaemon) is currently being used as part of development process to auto build project when any changes are made. The command to run this is `CompileDaemon -command="./dev-daily-journal go run main.go"` within the root directory of project
- With [TailwindCSS](https://tailwindcss.com/) being used. The following command will always be required to ensure my latest changes to the styles are put into effect; `./tailwindcss -i ./styles/input.css -o ./styles/output.css --watch`. Remember to run this command within the `/html` directory. 

## Notes
- I have been thinking about whether to use `HandleFunc` or `NewServeMux`. The only difference seems to be the use of `DefaultServeMux` which has got to do with how/when it is being used when/if other packages are involved.
- Reference: [This talks about how the `DefaultServeMux` is a global instance and that is bad when using a 3rd party package](https://stackoverflow.com/questions/54678816/using-handlefunc-on-http-vs-mux) and I am further convinced where performance may be worse and the nuance in getting information as [discussed](https://stackoverflow.com/questions/30063442/when-to-use-golangs-default-mux-versus-doing-your-own). Lastly, the global instance [will prevent us from determining the request type and makes it difficult to acquire information about URL like ID](https://perennialsky.medium.com/understand-handle-handler-and-handlefunc-in-go-e2c3c9ecef03)
- I realize why external third parties are used instead of the in-built `NewServeMux`. An example is URL parameters or knowing method type of the request in the backend. There are ways of [accomplishing](https://stackoverflow.com/questions/29211241/go-url-parameters-mapping) this but I would like to try to keep my code as clean as possible. Nevertheless, cool to find out about this.

<br/><br/>

## About Golang (to be moved into project)
### Go init Function
- `func init()` is executed first prior to `func main()`. Can take advantage of this to [variables initialization, check/fix program's state, registering, running one-time computations, etc.](https://medium.com/golangspec/init-functions-in-go-eac191b3860a) but you have to be aware of the [order](https://tutorialedge.net/golang/the-go-init-function/) because if a function is declared prior to `func init()`, the function is initialized first or if another function is used from another package - that gets initialized first which we should not rely on. Remember, focus on writing systems that do not rely on order.

### Directory purposes and possible ways of structuring projects in Golang
- [Flat structure or layered architecture](https://blog.logrocket.com/flat-structure-vs-layered-architecture-structuring-your-go-app/) are 2 possible ways I would like to try when developing this project. The plan is to start with flat structure first and separate out concerns by layers (packages) later on
- It is important to note that there are slight variations too when applying the layered architecture structure. The first to define the [relevant interfaces within the layers themselves](https://dev.to/codypotter/layered-architectures-in-go-3cg8) without [defining more generalized / business related models and interfaces in the layer](https://www.gobeyond.dev/packages-as-layers/) where the inner layers would import from the [business/domain specific layer which can be found on the highest layer](https://github.com/benbjohnson/wtf)

<br/><br/>

## About Others
### HTML Button
- I find out that [a button will require a form owner](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/button) to actually peform an action to the specified end point; to better explain this. A form tag is required to encapsulate the button tag for the button to actually send a request to the specified endpoint (in its action attribute). Note; to send a request - can use anchor tags instead which by default sends a `GET` request to the server

<br/><b/>

## Mongo DB introduction (to be moved into project)
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

<details><summary><h4>MongoDB Atlas Search</h4></summary>
  <ol>
    <li><details><summary><h4>Using relevance based search and search Index MongoDB</h4></summary>
      <ul>
        <li>Relevance search is different from database search. It starts with search indexes which creates a reference for records to relevance-based search usage</li>
        <li>Database search is used by developers and system administrators to help make database queries more efficient. These are the components of a search index</li>
        <pre><code>
        {
          "analyzer": "lucene.standard",
          "searchAnalyzer": "lucene.standard",
          "mappings": {
            "dynamic": true
          }
        }
        </pre></code>
      </ul>
    </details></li>
    <li><details><summary><h4>Creating search Index with dynamic mapping in MongoDB</h4></summary>
      <ul>
        <li>Search index define how relevance based search should perform. Dynamic mapping would ensure all fields are indexed except for booleans, objectIds and timestamps</li>
        <li>If dynamic mapping is used, specific fields will not need to be specified for field mappings on Mongo Atlas UI</li>
        <li>It is also possible to dedicated weights to specific fields so that scoring would be properly distributed as expected</li>
        <li>An example of utilizing the newly created search index (dynamic field mapping) and searching for the relevance text of "search" across multiple fields for each document which we then assign relevance scoring acquired for each document based on the relevance text - 'notepad'</li>
        <pre><code>
        db.sales.aggregate([
          {
            $search: {
              index: 'sample_supplies-sales-dynamic',
              text: {
                query: 'notepad', path: { 'wildcard': '*' }
              } } },
          {
            $set: {
              score: { $meta: "searchScore" }
              }
          }
        ])
        </pre></code>
      </ul>
    </details></li>
    <li><details><summary><h4>Creating search Index with static field mapping in MongoDB</h4></summary>
      <ul>
        <li>Static indexing on specific fields makes query quicker, saves processing power and time since we only focus on specific fields to be indexed</li>
      </ul>
    </details></li>
    <li><details><summary><h4>Using $search and Compound Operators in MongoDB</h4></summary>
      <ul>
        <li>Using compound operator, can assign different weights to specific fields to ensure we acquire most important results</li>
        <li>Will need to create a $search stage first (which will be created in the aggregation stage via Mongo Atlas). Within the $search stage, we will have a $compound component where we specific the weight of each field with commands like <code>must</code>(field must have specific value), <code>must not</code>, <code>filter</code> and <code>should</code> (where we assign weights)</li>
        <li>Example of creating a $search stage within the aggregation stage where the text "field" must be found in the habitat fieldd and a weight of 5 is allocated to wingspam_cm values that are greate than 45 in any documents.</li>
        <pre><code>
        {
          "compound": {
            "must": [{
              "text": {
                "query": "field",
                "path": "habitat"
              }
            }],
            "should": [{
              "range": {
                "gte": 45,
                "path": "wingspan_cm",
                "score": {"constant": {"value": 5}}
              }
            }]
          }
        }
        </pre></code>
        <li>Note that <code>filter</code> will not affect $search stage</li>
        <li>Additional examples. This aggregation stage has a $search stage within it that filters out documents that does not have the query text Online" within the purchaseMethod field and it then assigns a constant weight to the items sub document name field that contains the query text notepad in it. Finally, we project specicfic fields to cover the query and project a new field called score and acquire the aggregated score for documents found or determined to be relevant</li>
        <pre></code>
        db.sales.aggregate([
          {
            $search: 
            {
              index: 'sample_supplies-sales-dynamic',
              "compound": 
              {
                "filter": [{
                    "text": {"query": "Online", "path": "purchaseMethod"}
                  }],
                "should": [{
                    "text": {
                      "query": "notepad",
                      "path": "items.name",
                      "score": { "constant": { "value": 5 } }
                    }]
              }
            }
          },
          {
            $project: { "items.name": 1, "purchaseMethod": 1, "score": {$meta: "searchScore"} }
          }
        ])
        </pre></code>
      </ul>
    </details></li>
    <li><details><summary><h4>Group search results by using facets in MongoDB</h4></summary>
      <ul>
        <li>Allows us to categorize or group search results. These can be found in the $searchMeta and it is located there because it indicates how results have been aggregated and outputted.</li>
        <li>We then categorize results or place them in "buckets" with the use of $facets</li>
      </ul>
    </li></details>
  </ol>
</details>

<details><summary><h4>MongoDB Data Modeling</h4></summary>
  <ol>
    <li><details><summary><h4>Introduction to Data Modeling MongoDB</h4></summary>
      <ul>
        <li>It is how data will be stored and how the data will be related to one and another</li>
        <li>Good data modeling will ensure cost is saved, queries efficiency and identifying/predicting frequent document data access</li>
        <li>Note that each document can be different since MongoDB ensures that documents are flexible (polymorphism)</li>
        <li>Embedding a document would mean that we're correlated two documents with one and another</li>
      </ul>
    </details></li>
    <li><details><summary><h4>Types of Data Relationships MongoDB</h4></summary>
      <ul>
        <li>One to one, one to many and many to many. To model relationships - we can reference or embed data models</li>
        <li>Note - data that is accessed together should be stored together. If data stored in multiple collections, database must scan through multiple different collections which increases cost</li>
        <li>Embedding is taking related data and insert into document while referencing is when we refer documents in another collection to our document.</li>
        <li>Example of embedding one to many would include information about other documents within an array for a single document</li>
        <li>Referencing would include object ids from documents that belong to a different collection</li>
      </ul>
    </details></li>
    <li><details><summary><h4>Modeling Data Relationships MongoDB</h4></summary>
      <ul>
        <li>Many ways to model a relationship for documents. Note how context may be lost when different approaches have been taken to implement the relationships between the documents</li>
      </ul>
    </details></li>
    <li><details><summary><h4>Embedding Data in Documents MongoDB</h4></summary>
      <ul>
        <li>Used for one to many and many to many relationships between documents. Useful because it makes use of data that are accessed together - should be stored together.</li>
        <li>Allow all data to be accessed in one place and allow dev to get and update a document in one query</li>
        <li>However, documents may grow too large causing latency and high memory usage. As documents are read in full to memory - which is detrimental to application performance</li>
        <li>To counter this, read more about schema anti pattern (avoid unbounded documents)</li>
      </ul>
    </details></li>
    <li><details><summary><h4>Referencing Data in Documents MongoDB</h4></summary>
      <ul>
        <li>Saving _id of another document (can be from the same collection or different) as a field value in a document</li>
      </ul>
    </details></li>
    <li><details><summary><h4>Scaling a Data Model MongoDB</h4></summary>
      <ul>
        <li>Ensuring scalability can achieve efficient query result times, memory usage, cpu usage and storage usage</li>
        <li>Avoid unbounded documents, think of how document may grow larger (example is embedding comments to a blog post and if comments continue to grow - so does the size of document - causing performance issue when writing/reading)</li>
      </ul>
    </details></li>
    <li><details><summary><h4>Using Atlas Tools for Schema Help MongoDB</h4></summary>
      <ul>
        <li>Take notes of schemma design patterns to follow. Avoid anti patterns like massive arrays, massive number of collections, bloated documents, unnecessary indexes, queries without indexes and data that is accessed together from multiple collections</li>
        <li>Data explorer and performance advisor can be used to identify these anti patterns</li>
      </ul>
    </details></li>
  </ol>
</details>

<details><summary><h4>MongoDB Transactions</h4></summary>
  <ol>
    <li><details><summary><h4>Introduction to ACID Transactions in MongoDB</h4></summary>
      <ul>
        <li>Group of database operations that will be completed as a unit or not at all</li>
        <li>Atomicity, consistency, isolation and durability</li>
        <li>Atomicity - all operations succeed or none at all and Consistency is changes made by operations are consistent with database constraints</li>
        <li>Isolations - multiple transactions occuring at once should not affect other occuring transactions and Durability - changes made should persist</li>
      </ul>
    </details></li>
    <li><details><summary><h4>ACID Transactions in MongoDB</h4></summary>
      <ul>
        <li>Determining single or multiple documents ACID transaction. In MongoDB, single documents are inherently ACID transaction</li>
        <li>Multi documents operations are inherently not atomic. Will require additional steps. Will need to ensure that it is a multi-document ACID transaction.</li>
        <li>Ensure understand requirements of app. MongoDB will 'lock' resources involved in a transaction which incrus cost. We will need to insure operations are wrapped in multi-document ACID transaction.</li>
      </ul>
    </details></li>
    <li><details><summary><h4>Using Transactions in MongoDB</h4></summary>
      <ul>
        <li>Determining single or multiple documents ACID transaction. In MongoDB, single documents are inherently ACID transaction</li>
        <li>Session is a group of database operations that are related to each other and should be run together</li>
        <li>Transaction has maximum runtime of less than one minute. Error like <code>MongoServerError: Transaction has been aborted</code> would mean that query would need to be run again</li>
        <li>An example of using a multi-document transaction to ensure it is an ACID transaction</li>
        <pre><code>
        const session = db.getMongo().startSession()
        session.startTransaction()
        const account = session.getDatabase('< add database name here>').getCollection('<add collection name here>')
        //Add database operations like .updateOne() here
        session.commitTransaction() // or session.abortTransaction() to abort and rollback transactions
        </pre></code>
      </ul>
    </details></li>
  </ol>
</details>
