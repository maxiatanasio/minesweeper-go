import MineswepperClient from "./minesweeper";

//example of usage of the library

(async function init() {
    // 1) Create a client

    const baseUrl = 'https://warm-meadow-68089.herokuapp.com/' // This can be obtain by envars for example
    const client = new MineswepperClient(baseUrl)

    let gameStatus;
    let gameUuid;

    // 2) Creating a game of size 10 x 10
    gameUuid = await client.create(10, 10);
    console.log({gameUuid})

    // 3) Getting game current status
    gameStatus = await client.getStatus(gameUuid);
    console.log({gameStatus});

    // 4) Clicking a cell
    gameStatus = await client.click(gameUuid, 4, 5);
    console.log({gameStatus});

    // 5) Flagging a cell
    gameStatus = await client.flag(gameUuid, 6, 7);
    console.log({gameStatus});
})();

