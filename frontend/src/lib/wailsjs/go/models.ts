export namespace config {
	
	export class ConfigUser {
	    nick: string;
	    user: string;
	    real_name: string;
	
	    static createFrom(source: any = {}) {
	        return new ConfigUser(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.nick = source["nick"];
	        this.user = source["user"];
	        this.real_name = source["real_name"];
	    }
	}
	export class ConfigChannel {
	    name: string;
	    password: string;
	
	    static createFrom(source: any = {}) {
	        return new ConfigChannel(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.password = source["password"];
	    }
	}
	export class ConfigServer {
	    id: string;
	    name: string;
	    host: string;
	    port: number;
	    ssl: boolean;
	    tls_skip_verify: boolean;
	    disable_sts: boolean;
	    password: string;
	    last_connected: boolean;
	    channels: ConfigChannel[];
	    user: ConfigUser;
	
	    static createFrom(source: any = {}) {
	        return new ConfigServer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.host = source["host"];
	        this.port = source["port"];
	        this.ssl = source["ssl"];
	        this.tls_skip_verify = source["tls_skip_verify"];
	        this.disable_sts = source["disable_sts"];
	        this.password = source["password"];
	        this.last_connected = source["last_connected"];
	        this.channels = this.convertValues(source["channels"], ConfigChannel);
	        this.user = this.convertValues(source["user"], ConfigUser);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace girc {
	
	export class Channel {
	    name: string;
	    topic: string;
	    user_list: string[];
	    // Go type: time
	    joined: any;
	    // Go type: CModes
	    modes: any;
	
	    static createFrom(source: any = {}) {
	        return new Channel(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.topic = source["topic"];
	        this.user_list = source["user_list"];
	        this.joined = this.convertValues(source["joined"], null);
	        this.modes = this.convertValues(source["modes"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Source {
	    name: string;
	    ident: string;
	    host: string;
	
	    static createFrom(source: any = {}) {
	        return new Source(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.ident = source["ident"];
	        this.host = source["host"];
	    }
	}
	export class Event {
	    source?: Source;
	    tags: {[key: string]: string};
	    // Go type: time
	    timestamp: any;
	    command: string;
	    params: string[];
	    sensitive: boolean;
	    echo: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Event(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.source = this.convertValues(source["source"], Source);
	        this.tags = source["tags"];
	        this.timestamp = this.convertValues(source["timestamp"], null);
	        this.command = source["command"];
	        this.params = source["params"];
	        this.sensitive = source["sensitive"];
	        this.echo = source["echo"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class User {
	    nick: string;
	    ident: string;
	    host: string;
	    channels: string[];
	    // Go type: time
	    first_seen: any;
	    // Go type: time
	    last_active: any;
	    // Go type: UserPerms
	    perms?: any;
	    // Go type: struct { Name string "json:\"name\""; Account string "json:\"account\""; Away string "json:\"away\"" }
	    extras: any;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.nick = source["nick"];
	        this.ident = source["ident"];
	        this.host = source["host"];
	        this.channels = source["channels"];
	        this.first_seen = this.convertValues(source["first_seen"], null);
	        this.last_active = this.convertValues(source["last_active"], null);
	        this.perms = this.convertValues(source["perms"], null);
	        this.extras = this.convertValues(source["extras"], Object);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

