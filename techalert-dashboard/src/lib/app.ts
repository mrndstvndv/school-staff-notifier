import { LazyStore } from "@tauri-apps/plugin-store";
import Database from "@tauri-apps/plugin-sql";

export const store: LazyStore = new LazyStore("settings.json");
let dbInstance: null | Database = null;
export const getDbInstance = async (host: string): Promise<Database> => {
	if (!dbInstance) {
		dbInstance = await Database.load(`mysql://root:@${host}/school`);
	}
	return dbInstance;
};

export const selectFromTable = async (db: Database, columns: string[], table: string) => {
	const columnsString = columns.join(', ');

	return await db?.select(
		`select ${columnsString} from ${table}`,
	);
}

export const insertToTable = async (db: Database, object: Object, table: string) => {

	let fields = Object.keys(object).map((item) => {
		return `\`${item}\``
	})

	let values = Object.values(object).map((item, index) => {
		return `'${item}'`
	})

	let query = `INSERT INTO \`${table}\` (${fields}) VALUES (${values})`

	await db.execute(query)
}
