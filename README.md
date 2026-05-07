Ты прав, глаз резануло. В примере с **Teal** логика была описана более декларативно (через аннотации БД), в то время как в **Lua** структура создавалась через функцию `blomma.struct`. 

Чтобы проект выглядел целостным, я привел их к единому знаменателю: теперь и там, и там мы используем аннотации/декораторы для описания схемы данных. Это подчеркивает, что Blomma — это **Data-First** фреймворк.

Вот исправленный и финальный вариант:

---

# 🌸 Blomma
> You write the logic. Blomma does the rest.

Blomma is an open-source framework designed for rapid backend development. With a high-performance core written in **Rust** and business logic powered by **Lua** or **Teal**, Blomma eliminates the "infrastructure tax." After running `blomma init`, your workspace contains only your code — zero boilerplate, zero manual wiring.

---

## The Problem
Starting a new project usually means spending days on repetitive setup instead of building the product:
- Auth & Identity management
- Database schema and connection pooling
- Inter-module communication architecture
- Sessions, migrations, and input validation

Blomma abstracts all of this away before you write your first line of business logic.

---

## Quick Start
```bash
blomma init
```
```
? Project name: my-app
? Language: Lua / Teal
? Use decorators: yes / no
? HTML preprocessor (JSX): yes / no
? Enable plugins:
  ✔ http
  ✔ postgres
  ✔ auth
✅ Project created. Structure ready. Happy coding!
```

---

## Unified Logic: Lua vs Teal

### 1. Modern Lua (With Decorators)
In Lua, you define your data schema and routes using clean, Python-style decorators. The Rust core parses these at startup to build your DB tables and API map.

```lua
@db.record
local Post = {
    title = "@type:string; @min:4; @max:50",
    body  = "@type:string; @max:2000"
}

@auth
@validate(Post)
@route("POST /posts")
function create_post(ctx)
    -- Logic is clean: get data, save to DB, return result
    local post = ctx:body()
    local user = ctx:user()
    return posts.create(post, user)
end
```

### 2. Type-Safe Backend (Teal)
In Teal, we use the same logic, but with the power of static typing and native attribute syntax. This allows for compile-time checks of your business logic.

```teal
@db.record
local record Post
    title: string @db.rule("length.min:4;length.max:50")
    body: string  @db.rule("length.max:2000")
end

@auth
@validate(Post)
@route("POST /posts")
function create_post(ctx: Context): Response
    local post = ctx:body() as Post
    local user = ctx:user()
    return posts.create(post, user)
end
```

---

## Powerful Extensions

### Built-in SSR (JSX Preprocessor)
Forget separate template engines. Write your UI directly in your logic files with native performance.

```lua
@route("GET /posts")
function index(ctx)
    local items = db.posts.all()
    return (
        <div class="posts">
            {for _, post in ipairs(items) do}
                <article>
                    <h1>{post.title}</h1>
                    <p>{post.body}</p>
                </article>
            {end}
        </div>
    )
end
```

### The Event Bus
Decouple your modules using a high-speed internal event bus.

```lua
-- 'Publications' module emits an event
blomma.emit("post.created", { post = post, author = user })

-- 'Notifications' module reacts
blomma.on("post.created", function(ctx, event)
    local subscribers = users.find_subscribers(event.author)
    notifications.send(subscribers, event.post)
end)
```

---

## Plugin System
The Blomma Core is a lightweight event-driven engine. All features are modular.

| Plugin | Responsibility |
|---|---|
| `blomma-http` | High-performance server, routing, middleware |
| `blomma-postgres` | Native ORM, auto-migrations, pooling |
| `blomma-auth` | Sessions, JWT, and `ctx:user()` helpers |
| `blomma-mailer` | SMTP / API Email delivery |
| `blomma-queue` | Persistent background job processing |

**Extensibility:** Override any standard behavior with a simple `override` call.
```lua
blomma.override("auth.register", function(ctx)
    -- Your custom registration logic here
end)
```

---

## Key Features
- **Hot Reload:** Instant logic updates without restarting the Rust core.
- **Fault Tolerance:** Isolated handlers ensure one module's crash won't kill the server.
- **Performance:** Leveraging `mlua` for near-native execution speeds.
- **Leak Detection:** Core-level monitoring for memory and event loop health.

---

## Status
🚧 **In Development.** Core is being built with Rust, leveraging [mlua](https://github.com/mlua-rs/mlua).

---
*blomma — Swedish for "flower". Because a good tool should simply bloom, not get in your way.*
