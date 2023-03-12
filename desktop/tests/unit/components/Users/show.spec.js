import { mount } from "@vue/test-utils";
import ViewPage from "@/components/Users/View.vue";
import { createStore } from "vuex";

describe("ViewUser.vue", () => {
  afterEach(() => {
    jest.restoreAllMocks();
  });

  const mockRoute = {
    params: {
      id: 1,
    },
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
      },
    });
  };

  function factory() {
    const store = createVuexStore();

    return mount(ViewPage, {
      global: {
        plugins: [store],
        mocks: {
          $route: mockRoute,
        },
      },
    });
  }

  it("renders props.msg when passed", async () => {
    const msg =
      "AÇÕES Post 1 Visualização de usuáriofake@gmail.comEditarRemover";
    const wrapper = factory();

    console.log(wrapper.emitted());

    expect(mockRoute.params.id).toBe(1);
    expect(wrapper.text()).toMatch(msg);
    expect(wrapper.html()).toMatchSnapshot();
  });
});
