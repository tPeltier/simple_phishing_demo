# Simple Phishing Demo

This is an extremely simplified demonstration of a phishing attack.
This demo was developed for a database security group presentation and as such
is limited in scope and "realism".

## Features

- A singular server with "good" and "evil" paths to simulate a real and attacker
  server.
- Two emails, a phishing and real one, that contain `localhost` links to the
  appropriate server paths.
- A `schema.sql` to set up a very simple example MySql database.
- A Bash script to create and run a MySql container.
- A Bash script to run the server.

> [!IMPORTANT]
> THIS DEMO IS FOR EDUCATIONAL PURPOSES ONLY AND SHOULD NOT BE USED FOR ANY
> OTHER INTENTION. This demonstration is neither fully functional nor "live" and
> should only be used for educational reasons. Any unauthorized use of this
> information, including but not limited to malicious activities, is
strictly prohibited and could result in legal action.
>
> The author(s) and contributors of this repository do not assume any
> responsibility or liability for the actions of users who may misuse this
> information. By accessing and using this material, you agree to
hold harmless the author(s) and contributors from any claims arising out of your
use of this demonstration.
>
> The contents of this repository are provided "as is" without warranty of any
> kind, either expressed or implied. In no event shall the authors be liable for
> any claim, damages or other liabilities, whether in an action of contract,
> tort or otherwise, arising from, out of or in connection with the use of this
> demonstration.

## Overview

This demonstration focuses on the implications of phishing emails, so it cuts
quite a few corners from real-world implementations:

- There are no real login checks for users.
- Both "servers" have immediate and total access to the database.
- All data is simulated or "default."
- The website is extremely minimal.
- The phishing and real emails are identical except for the links.

## Basic usage

- Ensure that a Docker runtime & Go are installed on your system.
- Clone repository & navigate into the demo's directory
- Run `docker_sql_sandbox.sh` to initialize and start a MySql container.
- Copy the contents of `schema.sql` into the MySql instance.
- Run `run_demo.sh` to initiate the server and connect it to the database.
- Open the emails in a email reader of your choice and examine the differences
  and effects of the different links.

> [!NOTE]
> This demonstration was developed on Linux for Linux.
> Running it on other platforms is up to the user and may require additional
> configuration.
