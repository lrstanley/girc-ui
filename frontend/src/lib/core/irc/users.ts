/**
 * Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
 * of this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { currentChannelName } from "$lib/core/irc/channels"
import { latestStateUpdate } from "$lib/core/irc/events"
import { currentServerID } from "$lib/core/irc/servers"
import { GetUsersForChannel } from "$wails/go/servermanager/ServerManager"
import { derived } from "svelte/store"

export interface Permissions {
  owner?: boolean
  admin?: boolean
  op?: boolean
  half_op?: boolean
  voice?: boolean
}

export interface UserPerms {
  channels?: Record<string, Permissions>
}

export interface Extras {
  name?: string
  account?: string
  away?: string
}

export interface User {
  nick?: string
  ident?: string
  host?: string
  channels?: string[]
  first_seen?: string
  last_active?: string
  perms?: UserPerms
  extras?: Extras
}

export const users = derived<
  [typeof currentServerID, typeof currentChannelName, typeof latestStateUpdate],
  User[]
>([currentServerID, currentChannelName, latestStateUpdate], ([$serverID, $channelName, $uid], set) => {
  if (!$serverID || !$channelName) return set([])

  GetUsersForChannel($serverID, $channelName).then((users) =>
    Array.isArray(users) ? set(users) : set([])
  )
})
