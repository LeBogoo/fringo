# Fringo

Did you ever want to make late-night rambling feel like a competitive sport? Then you’re ready for Fringo!

It’s bingo, but better. You cross off the things that happened on Stream.

## User stories

> "The moment he sneezed like a dad and followed it up with a 'Minecraft villager noise,' I knew tonight was special."- Bob

> "He started talking about 'Tekilia,' got sidetracked into 'Shrimp farm,' and capped it off with a 'terrible pick-up line.' Triple threat Fringo." - Also Bob

> "It only took 'Rages at game,' 'Laughing at his own joke,' and a solid 'Goblin laugh' before I won Fringo in under ten minutes." - Frank

> "When he said 'Do you think owls would survive under capitalism?' while sipping tea, I had to mark four squares at once." - Bobby

## How to set up

### Requirements

- Node.js
- Go

### Frontend

1. Run `cd frontend`
2. Run `npm install`
3. Run `npm run dev`

### Backend

1. Run `cd backend`
2. Run `go run .`

## How to build

### Requirements

- Docker

1. Run `docker build --build-arg VERSION=0.0.0 --build-arg GIT_COMMIT=Unknown --build-arg GH_REPO=lebogoo/fringo-themes -t fringo .`
2. Run `docker run -p 8080:8080 fringo`
