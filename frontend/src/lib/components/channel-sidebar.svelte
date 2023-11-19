<script lang="ts">
    import ProfileBottombar from "$lib/components/profile-bottombar.svelte"
    import { channels, currentChannel, currentChannelName } from "$lib/core/irc/channels"
</script>

<div class="flex flex-col justify-between flex-none w-56 bg-gray-800">
    <div class="overflow-y-auto text-sm">
        <ul class="px-2 py-3">
            {#each $channels as channel}
                <li
                    class="px-2 text-gray-500 rounded-sm hover:text-gray-200 hover:bg-gray-900"
                    class:bg-gray-750={channel.name === $currentChannel?.name}
                    class:!text-gray-200={channel.name === $currentChannel?.name}
                >
                    <a
                        href="#"
                        class="flex items-center"
                        on:click={currentChannelName.set(channel.name)}
                    >
                        <!-- <span class="text-xl">#</span> -->
                        <span class="ml-1">{channel.name}</span>
                    </a>
                </li>
            {/each}
        </ul>
    </div>
    <ProfileBottombar />
</div>
