/**
 * Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
 * of this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { GetUsersForChannel } from "$wails/go/main/App.js"
import { EventsOn } from "$wails/runtime/runtime.js"
import { onMount } from "svelte"
import { readable } from "svelte/store"

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

export const users = readable<User[]>([], (set) => {
  const updateUsers = (u: User[]) => {
    Array.isArray(u) ? set(u) : set([])
  }

  onMount(() => {
    GetUsersForChannel("#girc-ui").then(updateUsers)
    return EventsOn("irc-update-state", () => {
      GetUsersForChannel("#girc-ui").then(updateUsers)
    })
  })
})
