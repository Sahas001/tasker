export type User = {
	id: number;
	name: string;
	email: string;
	password: string;
}

export type TaskInputProps = {
	title: string;
	description: string;
	category: string;
	status: boolean;
	createdAt: string;
}


export type Task = {
	id: number;
	userId: number;
	title: string;
	description: string;
	category: string;
	status: boolean;
	createdAt: string;
	updatedAt: string;
}
