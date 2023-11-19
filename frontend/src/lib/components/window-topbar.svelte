<script lang="ts">
    import IconArrowDown from "~icons/fluent/arrow-down-16-regular"
    import IconBell from "~icons/fluent/alert-16-regular"
    import IconPin from "~icons/fluent/pin-16-regular"
    import IconAt from "~icons/fluent/mention-16-regular"
    import IconQuestionMark from "~icons/fluent/question-16-regular"
    import { currentServer } from "$lib/core/irc/servers"
    import { currentChannel } from "$lib/core/irc/channels"
</script>

{#if $currentServer}
    <div class="flex h-12 text-white">
        <div
            class="flex items-center justify-between flex-none w-56 px-3 py-2 bg-gray-800 border-b border-gray-900"
        >
            <div>{$currentServer.name || $currentServer.id}</div>
            <button>
                <IconArrowDown />
            </button>
        </div>
        <div
            class="flex items-center justify-between flex-1 px-4 border-b border-gray-900 bg-gray-750"
        >
            {#if $currentChannel}
                <div class="flex items-center">
                    <div class="ml-2 text-sm text-white">{$currentChannel.name}</div>
                    <div class="pl-3 ml-3 text-xs text-gray-400 border-l border-gray-600">
                        {$currentChannel.topic || "No topic set"}
                    </div>
                </div>
            {/if}
            <div class="flex items-center">
                <IconBell />
                <IconPin />
                <a href="#" class="ml-4 mr-4">
                    <form action="#" class="relative">
                        <input
                            type="text"
                            placeholder="Search"
                            class="px-2 py-1 text-xs text-gray-200 bg-gray-900 rounded"
                        />
                        <span class="absolute top-0 right-0 mr-1" style="top: 6px;">
                            <svg
                                class="w-4 h-4 text-gray-500 hover:text-gray-200"
                                fill="currentColor"
                                viewBox="0 0 24 24"
                                width="24"
                                height="24"
                            >
                                <path
                                    class="heroicon-ui"
                                    d="M16.32 14.9l5.39 5.4a1 1 0 0 1-1.42 1.4l-5.38-5.38a8 8 0 1 1 1.41-1.41zM10 16a6 6 0 1 0 0-12 6 6 0 0 0 0 12z"
                                />
                            </svg>
                        </span>
                    </form>
                </a>
                <IconAt />
                <IconQuestionMark />
            </div>
        </div>
    </div>
{/if}
