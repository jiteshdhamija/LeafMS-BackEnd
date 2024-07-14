// MongoDB Playground
// Use Ctrl+Space inside a snippet or a string literal to trigger completions.

// The current database to use.
use('test');

// Create a new document in the collection.
db.getCollection('employees').insertOne({
    "username": "admin",
    "password": "GoldenRatio16",
    "name": "Admin",
    "team": "God",
    "Designation": "Admin",
    "Approver": "N/A"
});
