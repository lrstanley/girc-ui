/**
 * Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
 * of this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { EventsOn } from "$wails/runtime/runtime.js"
import { writable } from "svelte/store"

export interface EventSource {
  name: string
  ident: string
  host: string
}

export interface Event {
  source?: EventSource
  tags?: Record<string, string>
  timestamp: string
  command: string
  params: string[]
  sensitive: boolean
  echo: boolean
}

// export const events = readable<Event[]>([], (set) => {
//   let events: Event[] = []

//   onMount(() => {
//     const dereg = EventsOn("IRC_ALL", (data) => {
//       console.log(data)

//       events.push(data)
//       set(events)
//     })
//     onDestroy(dereg)
//   })
// })

export const events = writable<Event[]>([])

export function eventSetup(): () => void {
  return EventsOn("IRC_ALL", (data) => {
    events.update((events) => {
      events.push(data)
      return events
    })
  })
}
