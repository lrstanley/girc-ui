/**
 * Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
 * of this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { EventsOn } from "$wails/runtime/runtime"
import { onMount } from "svelte"
import { readable } from "svelte/store"

export interface EventSource {
  name?: string
  ident?: string
  host?: string
}

export interface Event {
  source?: EventSource
  tags?: Record<string, string>
  timestamp?: string
  command?: string
  params?: string[]
  sensitive?: boolean
  echo?: boolean
}

export const latestStateUpdate = readable<number>(-1, (set) => {
  onMount(() => {
    return EventsOn("irc-update-state", (e: Event) => set(Date.now()))
  })
})

export const latestGeneralUpdate = readable<number>(-1, (set) => {
  onMount(() => {
    return EventsOn("irc-update-general", (e: Event) => set(Date.now()))
  })
})

// export const events = readable<Event[]>([], (set) => {
//   const events: Event[] = []

//   const updateEvents = (e: Event) => {
//     if (e == null) return

//     events.push(e)
//     set(events)
//   }

//   onMount(() => {
//     // TODO: fetch events that have already happened, from the backend.
//     return EventsOn("IRC_ALL", updateEvents)
//   })
// })
