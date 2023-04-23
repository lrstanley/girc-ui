export namespace girc {
	
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

