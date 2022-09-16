# Blomma
Blomma is a web server written in Golang that uses Lua as its internal scripting language.

#### How should it work?

Blomma is a kind of analogue of Apache PHP. Actually works on a similar principle. When receiving an HTTP request, the server calls the execution of the file responsible for the path along which the request came.

# Here is a list of tasks that I would like to implement:

- [ ] Path processing system (Actually, the base is done, but I want to redo it a bit)
- [ ] Built-in MongoDB database driver support
- [ ] Built-in WebSocket Support
- [ ] Module/library system
- [ ] Gopherlua updates to the latest versions of Lua (I'm not sure if I can do it, but I'll try)
- [ ] Config


#### Goal

Goal for first working version: Make the scripting engine and server capable of REST API development

I will be glad to any suggestions on the project!
