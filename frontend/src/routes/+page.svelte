<script lang="ts">
    import { writable } from "svelte/store"
    import type { Writable } from "svelte/store"

    import { AppShell, AppBar, AppRail, AppRailTile } from "@skeletonlabs/skeleton"

    const storeValue: Writable<number> = writable(0)

    import { events } from "$lib/core/irc/events"
    import { users } from "$lib/core/irc/users"
</script>

<AppShell slotHeader="wails-draggable">
    <svelte:fragment slot="sidebarLeft">
        <div class="flex flex-auto flex-row w-[275px] h-full">
            <AppRail selected={storeValue} regionDefault="bg-slate-800 overflow-y-auto">
                <AppRailTile label="Server 1" value={0}>(icon)</AppRailTile>
                <AppRailTile label="Server 2" value={1}>(icon)</AppRailTile>
                <AppRailTile label="Server 3" value={2}>(icon)</AppRailTile>
            </AppRail>
            <div class="flex flex-col overflow-y-auto grow bg-slate-900">
                <nav class="py-2 list-nav">
                    <ul>
                        {#each [...Array(25).keys()] as item}
                            <li class="px-2 text-sm">
                                <a href={item < 1 ? "/" : "/about"}>
                                    <span class="flex-auto">#channel-{item}</span>
                                </a>
                            </li>
                        {/each}
                    </ul>
                </nav>
            </div>
        </div>
    </svelte:fragment>
    <svelte:fragment slot="sidebarRight">
        <div class="flex flex-auto flex-row w-[225px] h-full">
            <div class="flex flex-col overflow-y-auto grow bg-slate-900">
                <nav class="py-2 list-nav">
                    <ul>
                        {#each $users as user}
                            <li class="px-2 text-sm">
                                <a>
                                    <span class="flex-auto">{user.nick} ({user.extras.name})</span>
                                </a>
                            </li>
                        {/each}
                    </ul>
                </nav>
            </div>
        </div>
    </svelte:fragment>
    <svelte:fragment slot="pageHeader">
        <AppBar>
            <svelte:fragment slot="lead">(icon)</svelte:fragment>
            Channel subject
            <svelte:fragment slot="trail">(actions)</svelte:fragment>
        </AppBar>
    </svelte:fragment>
    <div class="flex flex-col flex-grow-0 flex-shrink max-h-[500px] overflow-y-auto">
        <!-- TODO -->
        <div class="flex flex-col flex-grow-0 flex-shrink max-h-[490px] overflow-y-auto">
            {#each $events as event}
                <div class="odd:bg-slate-900/70 even:bg-slate-700/70">
                    <span class="text-purple-600">@{event.source?.name}@{event.source?.host}</span>
                    <span>
                        {event.params?.join(" ")}
                    </span>
                </div>
            {/each}
        </div>
    </div>
    <svelte:fragment slot="pageFooter">
        <div class="p-2">
            <input class="rounded-sm input" type="text" placeholder="Input" />
        </div>
    </svelte:fragment>
</AppShell>
