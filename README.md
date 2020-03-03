# Gourd

#### A bot wrapper, command handler, and listener manager for the Disgord library

For a quick usage example, see the examples folder. For the moment, it covers creating a Gourd wrapper, and 
registering a module with commands. The wiki will be created soon(tm) and will cover all facets of Gourd in detail.

I do not have a dedicated discord server to support this project [yet?] 
but I can be found on the official unofficial DiscordAPI server, 
and reached via DM at Salmonllama#5727.

#### Important things to note:

- When creating commands, the first alias always becomes the command name
- If a command is used without aliases supplied, Gourd will provide an empty slice
- If a user is an administrator, it is only accounted for when using a PermissionInhibitor.

#### TODO:

- Proper Error Handling!!!! (critical priority)
- Finishing KeywordInhibition (low priority?)
- Built-in SQLite database (very high priority)
- Guild-based prefixes (very high priority)
- User/Guild blacklists (I-don't-even-know-if-this-is-needed priority)
- Option to include built-in modules. Would include: 
    - Help command (meh)
    - Prefix management (high)
    - blacklist management (see blacklist priority)
