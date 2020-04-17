// This is the API Client script for any JS project that wants
// to use the minesweeper API

import axios from 'axios'

export default class MineswepperClient {

    axiosClient = null;

    constructor(baseURL) {
        this.axiosClient = axios.create({
            baseURL,
            headers: {
                'Accept': 'application/json',
            }
        });
    }

    async create(sizeX, sizeY) {
        try {
            const response = await this.axiosClient.get('game/start/' + sizeX + '/' + sizeY);
            return response.data.uuid;
        } catch {
            throw new Error("There was an error creating the game");
        }
    }

    async getStatus(uuid) {
        try {
            const response = await this.axiosClient.get('game/status/' + uuid);
            return response.data.game;
        } catch {
            throw new Error("There was an error getting the game status");
        }
    }

    async click(uuid, x, y) {
        try {
            const response = await this.axiosClient.get(`game/click/${uuid}/${x}/${y}`)
            return response.data.game;
        } catch {
            throw new Error("There was an error doing a click");
        }
    }

    async flag(uuid, x, y) {
        try {
            const response = await this.axiosClient.get(`game/flag/${uuid}/${x}/${y}`)
            return response.data.game;
        } catch {
            throw new Error("There was an error processing the flag")
        }
    }

}