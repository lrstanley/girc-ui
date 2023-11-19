/**
 * Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
 * of this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { latestStateUpdate } from "$lib/core/irc/events"
import { currentServerID } from "$lib/core/irc/servers"
import type { girc } from "$wails/go/models"
import { GetServerChannel, GetServerChannels } from "$wails/go/servermanager/ServerManager"
import { onMount } from "svelte"
import { derived, get, writable } from "svelte/store"

export const channels = derived<[typeof currentServerID, typeof latestStateUpdate], girc.Channel[]>(
  [currentServerID, latestStateUpdate],
  ([$id, $uid], set) => {
    if (!$id) return set([])

    GetServerChannels($id).then((channel) => set(channel))
  }
)

export const currentChannelName = writable<string | null>(null, (set) => {
  onMount(async () => {
    const serverID = get(currentServerID)
    const channels = await GetServerChannels(serverID)

    if (channels?.length > 0) set(channels[0].name)

    return latestStateUpdate.subscribe(async () => {
      const serverID = get(currentServerID)
      const channels = await GetServerChannels(serverID)

      return currentChannelName.update((id) => {
        if (id != null || id != undefined) return id

        return channels?.length > 0 ? channels[0].name : null
      })
    })
  })
})

export const currentChannel = derived<[typeof currentServerID, typeof currentChannelName], girc.Channel>(
  [currentServerID, currentChannelName],
  ([$serverID, $channelName], set) => {
    if (!$serverID || !$channelName) return null

    GetServerChannel($serverID, $channelName).then((channel) => set(channel))
  }
)
