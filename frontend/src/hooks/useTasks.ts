import { useEffect, useReducer } from "react";
import { Task } from "../types/type";
import axios from "axios";

enum ActionType {
	LOADING,
	SUCCESS,
	FAILED,
}

type State = {
	loading: boolean;
	data: Task | null;
	error: any;
}

const initialState: State = {
	loading: false,
	data: null,
	error: null,
}

function reducer(state: State, action: { type: ActionType; payload?: any }) {
	switch (action.type) {
		case ActionType.LOADING:
			return { ...state, loading: true };
		case ActionType.SUCCESS:
			return { ...state, loading: false, data: action.payload };
		case ActionType.FAILED:
			return { ...state, loading: false, error: action.payload };
		default:
			return initialState;
	}

}

export function useTasks({ id }: { id: number }) {
	const [state, dispatch] = useReducer(reducer, initialState);

	useEffect(() => {
		fetchData();
	}, []);

	const fetchData = async () => {
		dispatch({ type: ActionType.LOADING });
		try {
			const response = await axios.get(`http://localhost:8080/tasks/user/${id}`);
			const data = response.data;
			dispatch({ type: ActionType.SUCCESS, payload: data });
		} catch (error) {
			dispatch({ type: ActionType.FAILED, payload: error });
		}
	};

	return {
		...state
	};
}

