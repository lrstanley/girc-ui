<script lang="ts">
    import IconTheme from "~icons/fluent/paint-bucket-16-regular"
    import IconChat from "~icons/fluent/chat-16-regular"
    import IconMinimize from "~icons/fluent/minimize-24-regular"
    import IconMaximize from "~icons/fluent/maximize-24-regular"
    import IconClose from "~icons/fluent/dismiss-24-regular"

    import { AppBar } from "@skeletonlabs/skeleton"
    import "./styles.css"
    import { baseTheme, themes, currentTheme } from "$lib/core/theme"

    import { onMount, onDestroy } from "svelte"
    import { Quit, WindowMinimise, WindowToggleMaximise } from "$wails/runtime/runtime.js"
    import { eventSetup } from "$lib/core/irc/events"
    import { userSetup } from "$lib/core/irc/users"

    let loaded = false
    onMount(() => (loaded = true))
    currentTheme.subscribe((value) => {
        if (!loaded) return

        document.body.setAttribute("data-theme", value)
    })

    onMount(() => {
        const eventDestroy = eventSetup()
        const userDestroy = userSetup()

        onDestroy(eventDestroy)
        onDestroy(userDestroy)
    })
</script>

<svelte:head>
    <script>
        if (!window.hasOwnProperty("wailsbindings")) {
            let wails_ipc = document.createElement("script")
            wails_ipc.setAttribute("src", "/wails/ipc.js")

            let wails_runtime = document.createElement("script")
            wails_runtime.setAttribute("src", "/wails/runtime.js")

            document.head.appendChild(wails_ipc)
            document.head.appendChild(wails_runtime)
        }
    </script>

    {@html `<style>${$currentTheme ? themes[$currentTheme] : ""}</style>`}
    {@html `<style>${baseTheme}</style>`}
</svelte:head>

<div class="flex flex-col w-full h-full">
    <header class="flex wails-draggable rounded-tl-md">
        <AppBar
            padding="px-2 p-[2px] w-full"
            gridColumns="grid-cols-3"
            slotTrail="place-content-end"
        >
            <svelte:fragment slot="lead">
                <div class="inline-flex flex-row items-center">
                    <IconChat class="mr-1 text-white" />
                    <span
                        class="font-medium text-transparent bg-gradient-to-br from-blue-500 to-cyan-300 bg-clip-text box-decoration-clone"
                    >
                        girc-ui
                    </span>
                </div>
            </svelte:fragment>

            <svelte:fragment slot="trail">
                <div class="inline-flex flex-row items-center">
                    <div
                        class="w-[200px] input-group input-group-divider grid-cols-[auto_1fr_auto] rounded-sm"
                    >
                        <div class="!px-2 input-group-shim">
                            <IconTheme />
                        </div>
                        <select
                            class="p-0 text-sm font-medium rounded-sm select"
                            bind:value={$currentTheme}
                        >
                            <option value="skeleton">Skeleton</option>
                            <option value="modern">Modern</option>
                            <option value="rocket">Rocket</option>
                            <option value="seafoam">Seafoam</option>
                            <option value="vintage">Vintage</option>
                            <option value="sahara">Sahara</option>
                            <option value="hamlindigo">Hamlin Digo</option>
                            <option value="crimson">Crimson</option>
                        </select>
                    </div>

                    <button class="px-2 hover:text-white" on:click={() => WindowMinimise()}>
                        <IconMinimize />
                    </button>

                    <button class="px-2 hover:text-white" on:click={() => WindowToggleMaximise()}>
                        <IconMaximize />
                    </button>

                    <button class="pl-2 hover:text-white" on:click={() => Quit()}>
                        <IconClose />
                    </button>
                </div>
            </svelte:fragment>
        </AppBar>
    </header>

    <main class="flex flex-col flex-grow overflow-y-auto bg-surface-900">
        <slot />
    </main>
</div>
