/**
 * Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
 * of this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { latestGeneralUpdate, latestStateUpdate } from "$lib/core/irc/events"
import type { config } from "$wails/go/models"
import { GetServerConfig, GetServersIDs } from "$wails/go/servermanager/ServerManager"
import { onMount } from "svelte"
import { derived, writable } from "svelte/store"

export const servers = derived<typeof latestGeneralUpdate, string[]>(
  latestGeneralUpdate,
  ($gid, set) => {
    GetServersIDs().then((ids) => (Array.isArray(ids) ? set(ids) : set([])))
  }
)

export const currentServerID = writable<string | null>(null, () => {
  onMount(async () => {
    const ids = await GetServersIDs()
    if (ids.length > 0) currentServerID.set(ids[0])

    return latestStateUpdate.subscribe(async () => {
      const ids = await GetServersIDs()

      return currentServerID.update((id) => {
        if (id != null || id != undefined) return id

        return ids.length > 0 ? ids[0] : null
      })
    })
  })
})

export const currentServer = derived<typeof currentServerID, config.ConfigServer | null>(
  currentServerID,
  ($id, set) => {
    if (!$id) return null

    GetServerConfig($id).then((config) => set(config))
  }
)
