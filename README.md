# CdA - Centre d’ambiance frontend

This application is the frontend for the CdA controller unit.

## Database Setup

The application is backed by a SQLite3 database. Edit "database.yml" to use the correct database file location.

### Create Your Databases

Create all databases (developement, test and production) with

	$ buffalo db create -a -d

and run all migrations

	$ buffalo db migrate

and seed database with defaults

	$ buffalo task db:seed

Alternatively 

	$ buffalo setup -d

can be executed what runs all three steps above and additionally setups the asset pipeline (`npm install` or `yarn install`) and runs all tests. 


## Starting the Application

Start the dev server with

	$ buffalo dev

and access the application at [http://127.0.0.1:3000](http://127.0.0.1:3000).

Starting the server on a specific port:

	$ PORT=3001 buffalo dev
	
