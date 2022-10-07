import asyncio
import os
import sqlite3
from time import sleep
from pyrogram import Client

api_id = os.environ["API_ID"]
api_hash = os.environ["API_HASH"]

con = sqlite3.connect("aneks/aneks.db")

async def main():
    channels = os.environ["CHANNELS"].split(",")

    async with Client("filldb", api_id, api_hash) as app:
        for channel in channels:
            print(f"start {channel}")

            async for message in app.get_chat_history(channel):
                if message.entities != None:
                    continue

                if message.media != None:
                    continue

                print(message.text)
                con.execute("INSERT INTO ANEKS VALUES (?);", (message.text,))
                con.commit()

                sleep(0.05)

            print(f"end {channel}")


asyncio.run(main())
conn.close()
