import { TimeEntry } from '@/businesslogic';
import { assign, Interpreter, Machine, EventObject, DoneInvokeEvent, send } from 'xstate';

const schema = {
  states: {
    Loading: {},
    Task: {
      Running: {},
      Stopped: {},
    },
    AddTask: {},
    AddingTask: {},
    StopTask: {},
    StoppingTask: {},
    Error: {},
  },
};

type Schema = typeof schema;

export interface Context {
  entry?: TimeEntry
}

export type Events =
  | { type: 'COVER_DIARY' }
  | { type: 'UNCOVER_DIARY' }
  | { type: 'PUBLISH' }
  | { type: 'SET_PUBLISH_TEXT', text: string }
  | { type: 'RELOAD' }
  | { type: 'HIDE_ERROR' };


export const machine = Machine<Context, Schema, EventObject>({
  id: 'currentTaskMachine',
  context: {},
  initial: 'Loading',
  on: {
    SAVE_TASK: 'AddingTask'
  },
  states: {
    Loading: {
      invoke: {
        src: 'loadLastTask',
        onDone: [
          { target: 'Task.Running', actions: 'assignEntry', cond: (_, { data }) => !(data as TimeEntry).endTime },
          { target: 'Task.Stopped', actions: 'assignEntry' }
        ],
        onError: { actions: 'logError', target: 'Error' }
      }
    }
  },
  Task: {
    states: {
      Running: {
        on: {
          STOP_TASK: 'StopTask',
          ADD_TASK: 'AddTask'
        }
      },
      Stopped: {
        on: {
          RESTART_TASK: { actions: send((ctx) => ({ type: 'SAVE_TASK', entry: ctx.entry })) }
        }
      }
    }
  },
  AddTask: {

  },

}, {
  actions: {
    setCovered: assign<Context>({
      covered: true,
    }),
    setUncovered: assign<Context>({
      covered: false,
    }),
    resetErrorText: assign<Context>({
      errorText: '',
    }),
    setPublishText: assign<Context>({
      publishText: (_, { text }: any) => text,
    }),
    clearPublishText: assign<Context>({
      publishText: '',
    }),
    assignResultToDiary: assign<Context>({
      diary: (_, evt) => (evt as DoneInvokeEvent<string>).data,
    }),
    assignErrorMessage: assign<Context>({
      errorText: (_, { data }: any) => data.toString(),
    }),
    // tslint:disable-next-line no-console
    logError: (_, { data }: any) => console.error(data.toString()),
  },
});

export type DiaryMachine = Interpreter<
  Context,
  Schema,
  Events
>;
