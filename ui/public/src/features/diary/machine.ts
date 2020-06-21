import { assign, Interpreter, Machine, EventObject, DoneInvokeEvent } from 'xstate';

const schema = {
  states: {
    idle: {},
    publishing: {},
    loading: {},
  },
};

type Schema = typeof schema;

export interface Context {
  diary: string;
  covered: boolean;
  errorText: string;
}

const contextInitial: Context = {
  diary: '',
  errorText: '',
  covered: true,
};

export type Events =
  | { type: 'COVER_DIARY' }
  | { type: 'UNCOVER_DIARY' }
  | { type: 'PUBLISH', entry: string }
  | { type: 'RELOAD' }
  | { type: 'HIDE_ERROR' };


export const machine = Machine<Context, Schema, EventObject>({
  id: 'diaryMachine',
  context: contextInitial,
  initial: 'idle',
  on: {
    COVER_DIARY: {
      actions: assign<Context>({
        covered: true,
      }),
    },
    UNCOVER_DIARY: {
      actions: assign<Context>({
        covered: false,
      }),
    },
  },
  states: {
    idle: {
      on: {
        PUBLISH: 'publishing',
        RELOAD: 'loading',
        HIDE_ERROR: {
          actions: 'resetErrorText',
        },
      },
    },
    publishing: {
      entry: 'resetErrorText',
      invoke: {
        src: async (_, evt) => {
          // tslint:disable no-console
          console.log(`Publishing "${(evt as any).entry}"`);
          return 'A publish result';
        },
        onDone: {
          target: 'idle',
          actions: 'assignResultToDiary',
        },
        onError: {
          target: 'idle',
          actions: ['assignErrorMessage', 'logError'],
        },
      },
    },
    loading: {
      entry: 'resetErrorText',
      invoke: {
        src: async (_, evt) => {
          // tslint:disable no-console
          console.log(`Loading`);
          return 'A load result';
        },
        onDone: {
          target: 'idle',
          actions: 'assignResultToDiary',
        },
        onError: {
          target: 'idle',
          actions: ['assignErrorMessage', 'logError'],
        },
      },
    },
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
    assignResultToDiary: assign<Context>({
      diary: (_, evt) => (evt as DoneInvokeEvent<string>).data,
    }),
    assignErrorMessage: assign<Context>({
      errorText: (_, { data }: any) => data.toString(),
    }),
    logError: (_, { data }: any) => console.error(data.toString()),
  },
});

export type DiaryMachine = Interpreter<
  Context,
  Schema,
  Events
>;
