import { mount } from "@vue/test-utils";
import RemovePage from "@/components/Users/Remove.vue";
import { createStore } from "vuex";

describe("RemoveUser.vue", () => {
  afterEach(() => {
    jest.restoreAllMocks();
  });

  window.axios = {
    get: jest
      .fn(() => {
        return {
          then: jest.fn(() => "your faked response"),
        };
      })
      .mockName("axiosget"),
    delete: jest
      .fn(() => {
        return {
          then: jest.fn(() => "your faked response"),
        };
      })
      .mockName("axiosPut"),
  };

  const mockRoute = {
    params: {
      id: 1,
    },
  };

  const mockRouter = {
    push: jest.fn(),
  };

  const createVuexStore = () => {
    return createStore({
      state() {
        return {
          Users: {
            oneUser: {
              ID: "1",
              name: "Post 1",
              email: "fake@gmail.com",
              username: "John Doe",
              password: "123",
            },
          },
        };
      },
      mutations: {
        updateUser(state, res) {
          state.oneUser = res;
        },
      },
      actions: {
        getUser(ctx, res) {
          ctx.commit("updateUser", res);
        },
        removeUser(ctx, res) {},
      },
    });
  };

  function factory() {
    const store = createVuexStore();

    return mount(RemovePage, {
      global: {
        plugins: [store],
        mocks: {
          axios: window.axios,
          $route: mockRoute,
          $router: mockRouter,
        },
      },
    });
  }

  it("renders props.msg when passed", async () => {
    const msg =
      "AÇÕES Post 1 Remoção de usuário Tem certeza que quer remover este usuário, essa ação não poderá ser desfeita! Não removerApagar definitivamente";
    const wrapper = factory();

    await wrapper.vm.remove();

    expect(wrapper.text()).toMatch(msg);

    await wrapper.findAll("a")[1].trigger("click");
    await wrapper.findAll("a")[0].trigger("click");
    console.log(wrapper.emitted());

    expect(wrapper.html()).toMatchSnapshot();
  });
});
