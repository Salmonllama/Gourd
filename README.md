# Gourd

#### A bot wrapper, command handler, and listener manager for the Disgord library

For a quick usage example, see the examples folder. For the moment, it covers creating a Gourd wrapper, and 
registering a module with commands.

I do not have a dedicated discord server to support this project [yet?] 
but I can be found on the official unofficial DiscordAPI server, 
and reached via DM at Salmonllama#5727.

#### Important things to note:

- When creating commands, the first alias always becomes the command name
- If a command is used without aliases supplied, Gourd will provide an empty slice

#### TODO:

- Permissions system:
    - Role-based restriction
    - Bot-owner restriction
    - PermissionsBit restriction
    - Specifically-assigned restrictions
        - Adding a custom permission keyword to a user
        - Restricting commands to custom permission keywords
- Built-in SQLite database
- Guild-based prefixes
- User/Guild blacklists
- Option to include built-in modules. Would include:
    - Help command
    - Prefix management
    - blacklist management
