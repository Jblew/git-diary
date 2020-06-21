import { assign, Interpreter, Machine, EventObject } from 'xstate';

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
  | { type: 'RELOAD' };


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
        PUBLISH: 'pubishing',
        RELOAD: 'loading',
      },
    },
    publishing: {
      entry: 'resetErrorText',
    },
    loading: {
      entry: 'resetErrorText',
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
  },
});

export type CampsiteWizardMachineInterpret = Interpreter<
  Context,
  Schema,
  Events
>;
