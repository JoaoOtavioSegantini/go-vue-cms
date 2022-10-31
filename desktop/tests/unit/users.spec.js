import Users from "@/components/Users.vue";
import { mount } from "@vue/test-utils";
import { createStore } from "vuex";

describe("Users.vue", () => {
  const createVuexStore = () => {
    return createStore({
      state() {
        return {
          Users: {
            all: [{ ID: "1", name: "John Doe", email: "johndoe@gmail.com" }],
          },
        };
      },
      mutations: {
        updateUsersList(state) {
          state.Users = [];
        },
      },
      actions: {
        listUsers(ctx) {
          const res = mockGet();
          ctx.commit("updateUsersList", res);
        },
      },
    });
  };

  function factory() {
    const store = createVuexStore();
    return mount(Users, {
      global: {
        plugins: [store],
        mocks: {
          $route: {
            params: {
              id: '1'
            }
          }
        }
      },
    });
  }

  const mockGet = jest.fn();

  jest.mock("axios", () => ({
    get: () => mockGet(),
  }));

  it("renders props when passed", () => {
    const msg = "AÇÕES Usuários Gerenciamento de usuários#nameemail1John Doejohndoe@gmail.comver";

    const wrapper = factory();

    expect(mockGet).toHaveBeenCalled();
    expect(wrapper.text()).toMatch(msg);
    expect(wrapper.html()).toMatchSnapshot();
  });
});
