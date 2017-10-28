# Conway's Game Of Life, Event Sourced

Incomplete example from https://www.meetup.com/DDD-CQRS-ES/events/243795846/

## Run:

```docker-compose up```

## Try:

### "Play" Command:

```./api-playground/curls/game-cmd-play.sh```

### Grid Read Model:

```./api-playground/curls/gameoflife-r-grid.sh```

## See the events in the event store:

http://localhost:2113/web/index.html#/streams/Game-gameId-1

Log in with username "admin", password "changeit".

## API Docs:

http://localhost:3999/GameOfLife.article