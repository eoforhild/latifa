# LATIFA
## Inspiration
Signal recently revealed a protocol called the PQXDH which is a key exchange algorithm that promises security against theoretical quantum adversaries. We wanted to implement this protocol alongside a file sharing app focused on ephemeral file transfers without the server hosting the data knowing anything about the actual files.

## What it does
The Rust infrastructure can complete the first two step of the PQXDH protocol, though the last step of mutually authenticating another user with the current remains unimplemented due to time constraints. The Flutter app can communicate somewhat with the Rust backend to facilitate this communication with the Go server. At this point, the backend can technically upload an encrypted file as a shared secret key is already derived, but the app's implementation to actually use it is not developed. It's not that useful, however, as the third step is what actually allows the end user to decrypt this file.

## How we built it
We already have a bit of experience working and reading through Signal's reference material and their specific terminologies. We consulted the specifications document heavily to begin structuring the protocol. We then used Rust, Go, MongoDB, and Flutter with Dart using a Rust bridge to connect everything together.

## Challenges we ran into
- We ran into problems with the school WiFi either being down or unresponsive, which made testing a server-client implementation between computers impossible
- We were very unfamiliar with NoSQL, trying to approach our data structures in a typical relational way which runs contrary to how NoSQL is designed to work
- Finding a way to properly bridge Rust and Flutter was tricky so testing the protocol was heavily dependent on the Flutter end working beforehand.
- The amount of variables we had to juggle in implementing the protocol was brain melting.
- Every part of the project had some kind of dependency on one another, so progression on individual parts heavily depends on everyone making progress. This meant that we had to constantly help each other out to walk through each of our thought processes.

## Accomplishments that we're proud of
We are proud of how much we learned and adapted to the various challenges presented by such a daunting project.

## What we learned
We learned a whole lot about how Flutter and Dart interacts with Rust. We also gained significant experience in how to even begin conceptualizing and structuring such a large project.

## What's next for Latifa
We want to actually flesh this out to use as a basis for the Double Ratchet algorithm also penned by Signal. We are following in their footsteps.
