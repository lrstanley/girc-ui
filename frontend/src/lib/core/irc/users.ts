import { GetUsersForChannel } from "$wails/go/main/App.js"
import { EventsOn } from "$wails/runtime/runtime.js"
import { writable } from "svelte/store"

export interface Permissions {
  owner: boolean
  admin: boolean
  op: boolean
  half_op: boolean
  voice: boolean
}

export interface UserPerms {
  channels: Record<string, Permissions>
}

export interface Extras {
  name: string
  account: string
  away: string
}

export interface User {
  nick: string
  ident: string
  host: string
  channels: string[]
  first_seen: string
  last_active: string
  perms?: UserPerms
  extras: Extras
}

// export const users = readable<User[]>([], (set) => {
//   // TODO: looks like this isn't invoked until HMR reloads the page.
//   onMount(() => {
//     GetUsersForChannel("#girc-ui").then((data) => {
//       set(data)
//     })

//     const dereg = EventsOn("irc-update-state", () => {
//       GetUsersForChannel("#girc-ui").then((data) => {
//         set(data)
//       })
//     })
//     onDestroy(dereg)
//   })
// })

export const users = writable<User[]>([])

export function userSetup(): () => void {
  // TODO: looks like this isn't invoked until HMR reloads the page.
  GetUsersForChannel("#girc-ui").then((data) => {
    users.set(data)
  })

  return EventsOn("irc-update-state", () => {
    GetUsersForChannel("#girc-ui").then((data) => {
      users.set(data)
    })
  })
}
