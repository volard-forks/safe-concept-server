# safe-concept-server

## concept of server
each authorized user owns an encrypted file with application data.

server implements DDOS protection, code, SQL injections and provides file storage.
one user - one encrypted file.

this concept supports the use of multiple applications.

it's assumed it will be used by standalone client applications that can storage user data in a single file. without application global or overlapped data.

## project layers architecture
![layers.png](./readme-assets/layers.png)

## project folder architecture
- server;
	- libs (common tools):
		- domain (lib is linked by domain zone);
			- index (export tools);
			- tools;
		- index (export tools);
	- middleware:
		- CORS;
		- DDOS-traffic-analyzer;
		- HTTP-requests-analyzer;
		- DB-provider;
		- FS-provider;
	- guards:
		- guard;
	- routes:
		- module:
			- route-name (DB);
				- controller (works with request);
				- service (works with data);
				- caching service;
				- DB_adapter (multiple DB supporting);
				- DB_worker (SQL query sender);
			- route-name (FS);
				- controller;
				- service;
				- caching (only get_file);
				- FS_worker;