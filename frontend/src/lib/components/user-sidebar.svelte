<script lang="ts">
    import { users } from "$lib/core/irc/users"
    import { currentServer } from "$lib/core/irc/servers"
</script>

<div class="flex-none w-56 py-3 pl-3 pr-2 overflow-y-auto bg-gray-800">
    <h3 class="mb-2 text-xs font-semibold tracking-wide text-gray-500 uppercase">admins</h3>
    <ul class="mb-6">
        <li
            class="flex items-center px-2 py-1 my-2 text-gray-500 hover:text-gray-200 hover:bg-gray-750"
        >
            <button class="flex items-center">
                <span>
                    <button>
                        <img
                            class="w-8 h-8 rounded-full"
                            src="https://avatars.githubusercontent.com/u/4323180"
                            alt="avatar"
                        />
                    </button>
                </span>
                <span class="ml-2 text-indigo-500">lrstanley ({$currentServer?.id})</span>
            </button>
        </li>
    </ul>
    <h3 class="mb-2 text-xs font-semibold tracking-wide text-gray-500 uppercase">
        Online — {$users.length}
    </h3>
    <ul class="m-1">
        {#each $users as user}
            <li
                class="px-2 py-1 my-2 text-gray-500 rounded-sm hover:text-gray-200 hover:bg-gray-750"
            >
                <button class="flex items-center">
                    <button class="items-center">
                        <img
                            class="w-6 h-6 rounded-full"
                            src="https://avatars3.githubusercontent.com/u/22605062?s=460&v=4"
                            alt="avatar"
                        />
                    </button>
                    <span class="m-2">{user.nick} ({user.extras.name})</span>
                </button>
            </li>
        {/each}
    </ul>
</div>
