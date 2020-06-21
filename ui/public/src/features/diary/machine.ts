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
  publishText: string;
}

const contextInitial: Context = {
  publishText: '',
  diary: '',
  errorText: '',
  covered: true,
};

export type Events =
  | { type: 'COVER_DIARY' }
  | { type: 'UNCOVER_DIARY' }
  | { type: 'PUBLISH' }
  | { type: 'SET_PUBLISH_TEXT', text: string }
  | { type: 'RELOAD' }
  | { type: 'HIDE_ERROR' };


export const machine = Machine<Context, Schema, EventObject>({
  id: 'diaryMachine',
  context: contextInitial,
  initial: 'idle',
  on: {
    COVER_DIARY: {
      actions: 'setCovered',
    },
    UNCOVER_DIARY: {
      actions: 'setUncovered',
    },
    SET_PUBLISH_TEXT: {
      actions: 'setPublishText',
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
        src: 'publishEntry',
        onDone: {
          target: 'idle',
          actions: ['assignResultToDiary', 'clearPublishText'],
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
        src: 'readDiary',
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
