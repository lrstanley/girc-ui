/**
 * Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
 * of this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { writable } from "svelte/store"
import skeletonBase from "@skeletonlabs/skeleton/styles/all.css?inline"
import crimson from "@skeletonlabs/skeleton/themes/theme-crimson.css?inline"
import hamlindigo from "@skeletonlabs/skeleton/themes/theme-hamlindigo.css?inline"
import modern from "@skeletonlabs/skeleton/themes/theme-modern.css?inline"
import rocket from "@skeletonlabs/skeleton/themes/theme-rocket.css?inline"
import sahara from "@skeletonlabs/skeleton/themes/theme-sahara.css?inline"
import seafoam from "@skeletonlabs/skeleton/themes/theme-seafoam.css?inline"
import skeleton from "@skeletonlabs/skeleton/themes/theme-skeleton.css?inline"
import vintage from "@skeletonlabs/skeleton/themes/theme-vintage.css?inline"

export const themes = {
  skeleton,
  modern,
  rocket,
  seafoam,
  vintage,
  sahara,
  hamlindigo,
  crimson
}

export const baseTheme = skeletonBase

export const defaultTheme = "crimson"
export const currentTheme = writable<keyof typeof themes>(defaultTheme)
